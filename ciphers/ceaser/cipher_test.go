package ceaser

import "testing"

func TestCeaserEncryptDecrypt(t *testing.T) {

	testString := "If he had anything confidential to say, he wrote it in cipher, that is, by so changing the order of the letters of the alphabet, that not a word could be made out."

	k := 0

	for k <= 100 {
		eString := Encrypt(testString, k)
		eString = Decrypt(eString, k)
		if eString != testString {
			t.Errorf("got [%s] wanted [%s] for k=%d\n", eString, testString, k)
		}
		k++
	}
}
