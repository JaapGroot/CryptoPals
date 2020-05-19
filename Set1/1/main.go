package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	HexInput := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	b64 := hex2base64(HexInput)
	fmt.Printf("%s", b64)
	fmt.Printf("\n")
}

func hex2base64(hx []byte) []byte {
	str := make([]byte, hex.DecodedLen(len(hx)))
	_, err := hex.Decode(str, hx)
	if err != nil {
		panic(err)
	}
	output := make([]byte, 4*(len(str)/3))
	base64.StdEncoding.Encode(output, str)

	return output
}
