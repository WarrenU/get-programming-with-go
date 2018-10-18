package main

import (
	"fmt"
	"math/rand"
)

const distance = 62100000
const secondsPerDay = 86400

func main() {
	fmt.Println("Spaceline        Days  Trip Type    Price")
	fmt.Println("=========================================")
	spacelines := [3]string{"Space Adventures", "SpaceX", "Virgin Galactic"}
	tripTypes := [2]string{"One-way", "Round-trip"}

	for i := 0; i < 10; i++ {
		spaceline := spacelines[rand.Intn(3)]
		speed := rand.Intn(15) + 16
		days := distance / (speed * secondsPerDay)
		tripType := tripTypes[rand.Intn(2)]
		price := 20.0 + speed
		if tripType == "Round-trip" {
			price *= 2
		}

		fmt.Printf("%-16v %4v  %-12v $%4v\n", spaceline, days, tripType, price)
	}
}
