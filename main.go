//Author: Jonathan Parlett
// Program Name: gocipher
// Created: 2021-December-30

package main

import (
	"fmt"
	"gcipher/ciphers/ceaser"
	"gcipher/utils"
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

	// s := "If he had anything confidential to say, he wrote it in cipher, that is, by so changing the order of the letters of the alphabet, that not a word could be made out."
	s := "jsjrdkfqqnslgfhpgwjfpymwtzlmnrrnsjsyqzhnzx"

	// fmt.Println("Before Encryption: ", s)

	// s = ceaser.Encrypt(s, 5)

	fmt.Println("After Encryption: ", s)

	s = ceaser.Decrypt(s, 5)

	fmt.Println("After Decryption: ", s)

	s = readFile("samplein.txt")

	// freqTable := utils.MakeCharFreqTable(s)

	// for i := 'A'; i <= 'Z'; i++ {
	// fmt.Printf("%c: %.2f\n", i, 100*freqTable[rune(i)])
	// }

	freqTable := utils.MakeBigramFreqTable(s)
	// fmt.Println(freqTable)
	for k, v := range freqTable {
		fmt.Printf("%s: %.2f\n", k, 100*v)
	}

}
