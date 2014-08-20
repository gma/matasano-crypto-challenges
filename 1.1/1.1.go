package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	hex_data := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	bytes, err := hex.DecodeString(hex_data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Base 64: %v\n", base64.StdEncoding.EncodeToString(bytes))
}
