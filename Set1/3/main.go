package main

import (
	"encoding/hex"
	"errors"
	"fmt"
)

func main() {
	input := HexDecode([]byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))

	fmt.Printf("%s\n", SingleByteXORCypher(input))
}

func SingleByteXORCypher(input []byte) []byte {
	for i := 0; i < 256; i++ {
		output := make([]byte, len(input))
		for j := 0; j < len(input); j++ {
			output[j] = input[j] ^ byte(i)
		}

		if IsValidString(output) {
			return output
		}
	}
	return nil
}

func IsValidString(s []byte) bool {
	var spaces int
	for _, r := range s {
		if r < 32 || r > 122 {
			return false
		}
		if r == 32 {
			spaces++
		}
	}
	if spaces < 3 {
		return false
	}
	return true
}

func HexDecode(input []byte) []byte {
	output := make([]byte, len(input)/2)
	_, err := hex.Decode(output, input)
	if err != nil {
		panic(err)
	}
	return output
}

func HexEncode(input []byte) []byte {
	output := make([]byte, len(input)*2)
	hex.Encode(output, input)
	return output
}

func FixedXOR(input1, input2 []byte) []byte {
	if len(input1) != len(input2) {
		panic(errors.New("Input not the same length!"))
	}
	output := make([]byte, len(input1))
	for i := 0; i < len(input1); i++ {
		output[i] = input1[i] ^ input2[i]
	}
	return output
}
