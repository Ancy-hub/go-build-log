package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader:=bufio.NewReader(os.Stdin)
	s,_:=reader.ReadString('\n')
	s = strings.ToUpper(s)
	vowels := "AEIOU"
	var vcount, ccount int
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			if strings.ContainsRune(vowels, char) {
				vcount++
			} else {
				ccount++
			}
		}
	}
	
	fmt.Printf("Vowels: %d, Consonants: %d", vcount, ccount)
}