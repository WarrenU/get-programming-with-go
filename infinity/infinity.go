package main

import (
	"fmt"
	"math/rand"
)

func main() {
	degrees := 0
	for {
		fmt.Println(degrees)
		degrees += 1
		if degrees >= 10 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}
