package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for countdown := 10; countdown >= 0; countdown-- {
		fmt.Println(countdown)
		time.Sleep(time.Second)
		if rand.Intn(100) == 3 {
			fmt.Println("Unsuccessful launch")
			break
		}

		if countdown == 0 {
			fmt.Println("Successful launch")
			fmt.Printf("%v is the value of countdown", countdown)
		}
	}
}
