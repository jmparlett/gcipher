package substitution

import (
	"math/rand"
	"strings"
	"time"
)

var asciiPerm [256]int

func initPerm(asciiPerm []int) {

	for i := 0; i < 256; i++ {
		asciiPerm[i] = i
	}

}

func swap(a *int, b *int) {
	//XOR swap
	*a ^= *b
	*b ^= *a
	*a ^= *b
}

func randomizePerm(asciiPerm []int, seed int64) {

	rand.Seed(seed)

	for k := len(asciiPerm) - 1; k > 1; k-- {
		swap(&asciiPerm[k], &asciiPerm[rand.Intn(k)])
	}

}

func genInversePerm(asciiPerm []int) []int {

	inversePerm := make([]int, len(asciiPerm))

	for i, v := range asciiPerm {
		inversePerm[v] = i
	}

	return inversePerm
}

func Encrypt(s string) (string, int64) {
	//accepts an ascii string and returns an ascii string of ciphertext

	var cipherText strings.Builder
	cipherText.Grow(len(s))

	//seed for real random
	rand.Seed(time.Now().UnixNano())

	key := rand.Int63() // gen random key that will be seed for everything else

	initPerm(asciiPerm[:])

	randomizePerm(asciiPerm[:], key)

	for _, c := range s {
		cipherText.WriteRune(rune(asciiPerm[int(c)]))
	}

	return cipherText.String(), key
}

func Decrypt(s string, key int64) string {
	//accepts a cipher text and a key and returns the plaintext

	var plainText strings.Builder
	plainText.Grow(len(s))

	initPerm(asciiPerm[:])

	randomizePerm(asciiPerm[:], key)

	inversePerm := genInversePerm(asciiPerm[:])

	for _, c := range s {
		plainText.WriteRune(rune(inversePerm[int(c)]))
	}

	return plainText.String()
}
