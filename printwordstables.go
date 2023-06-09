package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, char := range a {
		PrtRune(char)
	}
}

func PrtRune(s string) {
	str_rune := []rune(s)
	for _, word := range str_rune {
		z01.PrintRune(word)
	}
	z01.PrintRune('\n')
}
