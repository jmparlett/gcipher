package rot

import (
	"strings"
	"unicode"
)

//all rot algos are invertible and static so no need to worry about variable keys
func rotate(c rune, k int, start rune) rune {
	//char, shift amt, start char,
	return rune(int(start) + ((int(c) + k - int(start)) % (2 * k)))
}

func Rot5(s string) string {
	var es strings.Builder
	es.Grow(len(s))

	for _, v := range s {
		if unicode.IsDigit(v) {
			es.WriteRune(rotate(v, 5, '0'))
		} else {
			es.WriteRune(v)
		}
	}
	return es.String()
}

func Rot13(s string) string {
	var es strings.Builder
	es.Grow(len(s))

	for _, v := range s {
		if unicode.IsLetter(v) {
			var start rune
			if unicode.IsLower(v) {
				start = 'a'
			} else {
				start = 'A'
			}
			es.WriteRune(rotate(v, 13, start))
		} else {
			es.WriteRune(v)
		}
	}
	return es.String()
}

func Rot18(s string) string {
	//rot 18 is just rot5+rot13
	return (Rot13(Rot5(s)))
}
