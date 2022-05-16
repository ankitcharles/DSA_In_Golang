package main

import "fmt"

// VowelsFinder Interface Definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// FindVowels MyString implements VowelsFinder Interface
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}
func main() {
	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name
	fmt.Printf("Vowels are %c", v.FindVowels())
}
