package main

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		SingleByteXORCypher(HexDecode(line))

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// SingleByteXORCypher checks string with all possible bytes
func SingleByteXORCypher(input []byte) []byte {
	for i := 0; i < 256; i++ {
		output := make([]byte, len(input))
		for j := 0; j < len(input); j++ {
			output[j] = input[j] ^ byte(i)
		}

		if IsValidString(output) {
			fmt.Printf("%s", output)
		}
	}
	return nil
}

// IsValidString check if sentence is English
func IsValidString(s []byte) bool {
	stringLen := len(s)
	if s[0] < 'A' || s[0] > 'Z' {
		return false
	}

	if s[stringLen-1] != '\n' {
		return false
	}

	prevState := 0
	currState := 0

	for i := 1; i < stringLen; i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			currState = 0
		} else if s[i] == ' ' {
			currState = 1
		} else if s[i] >= 'a' && s[i] <= 'z' {
			currState = 2
		} else if s[i] == '.' || s[i] == '\n' {
			currState = 3
		}

		if prevState == currState && currState != 2 {
			return false
		}
		if prevState == 2 && currState == 0 {
			return false
		}

		if currState == 3 && prevState != 1 {
			return true
		}
	}
	return false
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

// HexEncode encode hext string
func HexEncode(input []byte) []byte {
	output := make([]byte, len(input)*2)
	hex.Encode(output, input)
	return output
}

// FixedXOR XOR 2 strings
func FixedXOR(input1, input2 []byte) []byte {
	if len(input1) != len(input2) {
		panic(errors.New("input not the same length"))
	}
	output := make([]byte, len(input1))
	for i := 0; i < len(input1); i++ {
		output[i] = input1[i] ^ input2[i]
	}
	return output
}
