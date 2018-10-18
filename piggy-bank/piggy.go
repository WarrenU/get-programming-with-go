package main

import (
	"fmt"
	"math/rand"
)

func main() {
	currency := [3]float64{0.05, 0.10, 0.25}
	piggybank := 0.0
	for piggybank < 20.00 {
		piggybank += currency[rand.Intn(3)]
		fmt.Printf("$%5.2f\n", piggybank)
	}
}
