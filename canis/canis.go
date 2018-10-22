// Package main implements a convesion of `distance` (between Earth and Canis Major Dwarf) to `light years`.
package main

import (
	"fmt"
)

func main() {
	const distance = 236e15
	const lightSpeed = 299792
	const secondsInDay = 86400
	const daysInYear = 365

	years := distance / lightSpeed / secondsInDay / daysInYear

	fmt.Printf("Canis Major Dwarf is %.0f light years away from Earth.", years)
}
