package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
)

// KEY Encrypt key
const KEY string = "asdf"

// KEYSIZE Max lenght of the key
const KEYSIZE int = 40

// NAvgSorter sorts planets by axis.
type NAvgSorter []Hamming

func (a NAvgSorter) Len() int           { return len(a) }
func (a NAvgSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NAvgSorter) Less(i, j int) bool { return a[i].NAvgHD < a[j].NAvgHD }

// Hamming struct to save hamming details of calculation
type Hamming struct {
	CalcKeySize int
	HD          float64
	AvgHD       float64
	NAvgHD      float64
}

func main() {
	var input []byte
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		input = append(input, line...)
	}
	B64Decode := make([]byte, (len(input)*3)/4)
	base64.StdEncoding.Decode(B64Decode, input)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	HamData := CalcMatasano(B64Decode)
	sort.Sort(NAvgSorter(HamData))
	ProbKeySize := HamData[0].CalcKeySize

	for _, k := range HamData {
		fmt.Printf("CalcKeySize: %d, HD: %.6g, \tAvgHD: %.5g, \tNAvgHD: %.9g\n", k.CalcKeySize, k.HD, k.AvgHD, k.NAvgHD)
	}

	chunks := SplitInBlocks(B64Decode, ProbKeySize)
	var Key []byte
	for _, chunk := range chunks {
		_, char := SingleByteXORCypher(chunk)
		Key = append(Key, char)
	}

	result := RepeatingKeyXOR(B64Decode, Key)
	fmt.Printf("%s\n", result)
}

// SplitInBlocks Split full text in smaller parts
func SplitInBlocks(input []byte, keysize int) [][]byte {
	var blocks, chunks [][]byte

	for i := 0; i < len(input)-keysize; i = i + keysize {
		blocks = append(blocks, input[i:i+keysize])
	}
	for j := 0; j < keysize; j++ {
		var chunk []byte
		for _, block := range blocks {
			chunk = append(chunk, block[j])
		}
		chunks = append(chunks, chunk)
	}

	return chunks
}

// CalcMatasano Calculate matasano for all posibilities
func CalcMatasano(input []byte) []Hamming {
	HamData := make([]Hamming, KEYSIZE-1)
	for j := 2; j <= KEYSIZE; j++ {
		timesCalculated := 0.0
		total := 0.0
		for i := 0; i < len(input); i = i + j {
			if i+j*2 > len(input) {
				break
			}
			hd, err := CalcHammingDistance(input[i:i+j], input[i+j:i+j*2])
			if err != nil {
				panic(err)
			}

			total += float64(hd)
			timesCalculated++
		}
		HamData[j-2].CalcKeySize = j
		HamData[j-2].HD = total
		HamData[j-2].AvgHD = total / timesCalculated
		HamData[j-2].NAvgHD = total / timesCalculated / float64(j)
	}
	return HamData
}

// CalcHammingDistance calculates the hamming distance between 2 strings
func CalcHammingDistance(str1, str2 []byte) (int, error) {
	var HD int
	if len(str1) != len(str2) {
		return 0, errors.New("a and b not the same length")
	}

	for Byte := 0; Byte < len(str1); Byte++ {
		for Bit := 0; Bit < 8; Bit++ {
			mask := byte(1 << Bit)
			if (str1[Byte] & mask) != (str2[Byte] & mask) {
				HD++
			}
		}
	}

	return HD, nil
}

// SingleCharXOR XOR single char
func SingleCharXOR(text []byte, char byte) ([]byte, float64) {
	output := make([]byte, len(text))
	frequency := 0.0
	for i, Byte := range text {
		output[i] = Byte ^ char
		frequency += calcAlphScore(Byte ^ char)
	}

	return output, frequency
}

// SingleByteXORCypher checks string with all possible bytes
func SingleByteXORCypher(input []byte) ([]byte, byte) {
	highestValue := 0.0
	bestPhrase := make([]byte, len(input))
	var keyChar byte
	for i := 0; i < 128; i++ {
		output, frequency := SingleCharXOR(input, byte(i))
		if highestValue < frequency {
			bestPhrase = output
			keyChar = byte(i)
			highestValue = frequency
		}
	}
	return bestPhrase, keyChar
}

// RepeatingKeyXOR Implementation to XOR with key
func RepeatingKeyXOR(input, Key []byte) []byte {
	var output []byte
	for i, r := range input {
		key := Key[i%len(Key)]

		output = append(output, r^key)
	}

	return output
}

// HexEncode encode hext string
func HexEncode(input []byte) []byte {
	output := make([]byte, len(input)*2)
	hex.Encode(output, input)
	return output
}

// HexDecode decode hex string
func HexDecode(input []byte) []byte {
	output := make([]byte, len(input)/2)
	_, err := hex.Decode(output, input)
	if err != nil {
		panic(err)
	}
	return output
}

// calcAlphScore calculate the language score
func calcAlphScore(char byte) float64 {
	if char > 64 && char < 90 {
		char += 32
	}

	frequency := map[byte]float64{
		'a': 8.167, 'b': 1.492, 'c': 2.782, 'd': 4.253, 'e': 12.702,
		'f': 2.228, 'g': 2.015, 'h': 6.094, 'i': 6.966, 'j': 0.153,
		'k': 0.772, 'l': 4.025, 'm': 2.406, 'n': 6.749, 'o': 7.507,
		'p': 1.929, 'q': 0.095, 'r': 5.987, 's': 6.327, 't': 9.056,
		'u': 2.758, 'v': 0.978, 'w': 2.360, 'x': 0.150, 'y': 1.974,
		'z': 0.074, ' ': 12.8,
	}

	return frequency[char]
}
