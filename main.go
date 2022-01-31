//Author: Jonathan Parlett
// Program Name: gocipher
// Created: 2021-December-30

package main

import (
	"fmt"
	"gcipher/ciphers/substitution"
	"log"
	"os"
)

func readFile(s string) string {
	//take in a file name and read to a string
	data, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(data)
}

func main() {

	s := "If he had anything confidential to say, he wrote it in cipher, that is, by so changing the order of the letters of the alphabet, that not a word could be made out."
	//
	fmt.Println("Before Encryption: ", s)

	s, key := substitution.Encrypt(s)

	fmt.Println("After Encryption: ", s)

	s = substitution.Decrypt(s, key)

	fmt.Println("After Decryption: ", s)

	// for _, v := range s {
	// fmt.Printf("%d ", int(v))
	// }
	//
	// fmt.Println()
	//
	// for _, v := range s {
	// fmt.Printf("%c", rune(int(v)))
	// }

	// substitution.Test()
}
