//Author: Jonathan Parlett
// Program Name: gocipher
// Created: 2021-December-30

package main

import (
	"fmt"
	"gcipher/ciphers/ceaser"
	"gcipher/ciphers/rot"
)

func main() {

	s := "If he had anything confidential to say, he wrote it in cipher, that is, by so changing the order of the letters of the alphabet, that not a word could be made out."

	fmt.Println("Before Encryption: ", s)

	s = ceaser.Encrypt(s, 13)

	fmt.Println("After Encryption: ", s)

	s = ceaser.Decrypt(s, 13)

	fmt.Println("After Decryption: ", s)

	rstring := "This is 55 a String 1234587 containing numbers for rot5"

	fmt.Println("Before Encryption: ", rstring)
	rstring = rot.Rot5(rstring)
	fmt.Println("After Encryption: ", rstring)
	rstring = rot.Rot5(rstring)
	fmt.Println("After Decryption: ", rstring)

	fmt.Println("Before Encryption: ", rstring)
	rstring = rot.Rot13(rstring)
	fmt.Println("After Encryption: ", rstring)
	rstring = rot.Rot13(rstring)
	fmt.Println("After Decryption: ", rstring)

	fmt.Println("Before Encryption: ", rstring)
	rstring = rot.Rot18(rstring)
	fmt.Println("After Encryption: ", rstring)
	rstring = rot.Rot18(rstring)
	fmt.Println("After Decryption: ", rstring)
}
