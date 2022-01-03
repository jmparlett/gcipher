package utils

import "unicode"

func MakeCharFreqTable(s string) map[rune]float64 {

	var table = make(map[rune]float64)

	charCount := 0.0

	for _, v := range s {
		if unicode.IsLetter(v) {
			table[v]++
			charCount += 1
		}
	}

	for k, _ := range table {
		table[k] /= charCount
	}
	return table
}

func MakeBigramFreqTable(s string) map[string]float64 {
	//count all substrings of len 2
	var table = make(map[string]float64)

	bigramCount := 0.0

	for i := 1; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) && unicode.IsLetter(rune(s[i-1])) {
			table[s[i-1:i+1]]++
			bigramCount++
		}
	}

	for k, _ := range table {
		table[k] /= bigramCount
	}

	return table
}
