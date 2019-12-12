package main

import  "fmt"

type intSlice []int

func main() {

	intList := intSlice {3,5,7,8,9,11,13,17,23,45,67,81,90,94}

	numberToSearch := 12

	result := -1

	if numberToSearch >= intList[0] && numberToSearch <= intList[len(intList)-1] {
		result = binary(intList ,0 , len(intList)-1 , numberToSearch)
	}

	if result == -1 {
		fmt.Println("Number not found")
	} else {
		fmt.Printf("Number #%v found at #%v position", numberToSearch , result)
	}


}

func binary(arr []int, left,right,num int) int {

	if right>=left {
		mid := (left+right) /2
		fmt.Printf("Mid at position: %v\n", mid)
		if arr[mid] == num {
			return mid
		}

		if arr[mid] > num {
			return binary(arr,left,mid-1,num)
		} else {
			return binary(arr,mid+1,right,num)
		}

	}
	return -1
}
















