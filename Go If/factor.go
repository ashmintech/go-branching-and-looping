package main

import "fmt"

type intSlice []int

func main() {

	intList := intSlice {24,46,65,86,97,123,241,376,651}
	
	factors := make(map[int]intSlice)

	for _,i := range intList {
		for j:=1 ; j<=i ; j++ {
			if i%j == 0 {
				factors[i] = append(factors[i],j)
			}
		}
	}

	for x,y :=range factors {
		fmt.Printf("\nFactors of %v:\t %v", x,y)
	}
	
	fmt.Print("\nPrime Numbers:\t")
	for x,y := range factors {
		if len(y) != 2 {
			continue
		}
		fmt.Printf("%v ",x)
	}

}
