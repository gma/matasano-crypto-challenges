package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	buffer1 := "1c0111001f010100061a024b53535009181c"
	buffer2 := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"

	hex1, _ := hex.DecodeString(buffer1)
	hex2, _ := hex.DecodeString(buffer2)

	bytes := make([]byte, len(hex1))

	for i := 0; i < len(hex1); i++ {
		bytes[i] = hex1[i] ^ hex2[i]
	}

	if expected == hex.EncodeToString(bytes) {
		fmt.Println("Success!")
	} else {
		fmt.Println("Fail :-(")
	}
}
