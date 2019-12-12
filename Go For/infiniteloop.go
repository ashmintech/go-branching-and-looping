package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	go findRandomNumber(rand.Intn(100))
	time.Sleep(5 * time.Second)

}

func findRandomNumber(randomNumber int) {

	count :=1
	numberFound := false

	for {
		number := rand.Intn(1000000000)
		if number == randomNumber {
			numberFound = true
			break;
		}
		count++;
	}

	if numberFound {
		fmt.Printf("Number #%v found after %v attempt(s)", randomNumber , count)
	}

}