package main

import (
	"fmt"
	"math/rand"
)

func generateDate() (int, int, int) {
	year := rand.Intn(2018) + 1
	leap := year%400 == 0 || year%4 == 0 && year%100 != 0
	month := rand.Intn(12) + 1
	days := 31
	switch month {
	case 2:
		days = 28
		if leap {
			days = 29
		}
	case 4, 6, 9, 11:
		days = 30
	}

	day := rand.Intn(days) + 1
	return year, month, day
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(generateDate())
	}
}
