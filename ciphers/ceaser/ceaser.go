package ceaser

import (
	"fmt"
	"strings"
	"unicode"
)

var rotBound int = 26

func computeShift(v rune, k int, d bool) rune { //because ceaser shifts aren't symmetrical we cant use the more general rot shift
	if d { //forward
		if unicode.IsLower(v) {
			return (rune(((int(v) - int(rune('a')) + k) % rotBound) + int(rune('a'))))
		} else {
			return (rune(((int(v) - int(rune('A')) + k) % rotBound) + int(rune('A'))))
		}
	} else { //back
		if unicode.IsLower(v) {
			return (rune(((int(v) - int(rune('a')) + rotBound - k) % rotBound) + int(rune('a'))))
		} else {
			return (rune(((int(v) - int(rune('A')) + rotBound - k) % rotBound) + int(rune('A'))))
		}
	}
}

func Encrypt(s string, k int) string {
	k = k % rotBound //normalize

	//string builder is more efficient
	var es strings.Builder
	es.Grow(len(s))

	for _, v := range s {
		if unicode.IsLetter(v) {
			es.WriteRune(computeShift(v, k, true)) //shift foward
		} else {
			es.WriteRune(v)
		}
	}
	return es.String()
}

func Decrypt(s string, k int) string {
	k = k % rotBound //normalize

	//string builder is more effecient
	var es strings.Builder
	es.Grow(len(s))

	for _, v := range s {
		if unicode.IsLetter(v) {
			es.WriteRune(computeShift(v, k, false)) //shift back
		} else {
			es.WriteRune(v)
		}
	}
	return es.String()
}

func BruteForce(s string) { //brute force a ceaser cipher encrypted string by trying all combinations
	i := 0
	for i <= 26 {
		fmt.Println(Decrypt(s, i))
		i++
	}
}
