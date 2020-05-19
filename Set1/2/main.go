package main

import (
	"encoding/hex"
	"errors"
	"fmt"
)

func main() {
	input1 := HexDecode([]byte("1c0111001f010100061a024b53535009181c"))
	input2 := HexDecode([]byte("686974207468652062756c6c277320657965"))

	str := HexEncode(FixedXOR(input1, input2))

	fmt.Printf("%s", str)
	fmt.Printf("\n")

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
