package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"unicode"
)

func main() {
	encrypted := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	encrypted_bytes, _ := hex.DecodeString(encrypted)

	plain_text := make([]byte, len(encrypted_bytes))

	for i := 0; i < 256; i++ {
		single_byte := make([]byte, 1)
		single_byte[0] = byte(i)
		key := bytes.Repeat(single_byte, len(encrypted_bytes))
		for j := 0; j < len(encrypted_bytes); j++ {
			plain_text[j] = encrypted_bytes[j] ^ key[j]
		}
		chars := countCharacters(string(plain_text))
		fmt.Printf("Number of letters: %v\n", chars)
	}
}

func countCharacters(text string) int {
	count := 0
	for i := 0; i < len(text); i++ {
		if unicode.IsLetter(rune(text[i])) {
			count++
		}
	}
	return count
}
