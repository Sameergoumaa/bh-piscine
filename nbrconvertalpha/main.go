package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	argumet := os.Args
	for i := 1; i <= len(argumet)-1; i++ {
		name := []rune(argumet[i])
		for t := 0; t < len(name); t++ {
			j := name[t]
			z01.PrintRune(letters[j])
		}
	}
}
