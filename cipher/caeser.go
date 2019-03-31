// Package main deciphers a quote from a Julius Caesar: L fdph, L vdz, L frqtxhuhg.
package main

import (
	"fmt"
)

func decipher(phrase string) {
	for _, c := range phrase {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			c = c - 3
		}
		fmt.Printf("%c", c)
	}
}

func main() {
	phrase := "L fdph, L vdz, L frqtxhuhg."
	decipher(phrase)
}
