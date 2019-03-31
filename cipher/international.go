// Package international ciphers a spanish phrase.
package main

import "fmt"

func cipherROT13(phrase string) {
	for _, c := range phrase {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			c = c + 13
			if c > 'z' || c > 'Z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
}

func main() {
	phrase := "Hola Estación Espacial Internacional"
	cipherROT13(phrase)
}
