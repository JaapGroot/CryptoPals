package main

import (
	"encoding/hex"
	"fmt"
)

// KEY Encrypt key
const KEY string = "ICE"

func main() {
	input := []byte("Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal")

	encodedMsg := HexEncode(RepeatingKeyXOR(input))
	fmt.Printf("%s\n", encodedMsg)
	decodedMsg := RepeatingKeyXOR(HexDecode(encodedMsg))
	fmt.Printf("%s\n", decodedMsg)
}

// RepeatingKeyXOR Implementation to XOR with key
func RepeatingKeyXOR(input []byte) []byte {
	var output []byte
	for i, r := range input {
		key := KEY[i%len(KEY)]

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
