package main

import "fmt"

func main() {
	numbers := []int32{10, 15, 3, 7}
	fmt.Println(searchAddup(&numbers, 17))
}

func searchAddup(numbers *[]int32, k int32) (bool, int32, int32) {
	fmt.Println(*numbers)
	for i := 0; i < len(*numbers); i++ {
		for j := i + 1; j < len(*numbers); j++ {
			if (*numbers)[i]+(*numbers)[j] == k {
				return true, (*numbers)[i], (*numbers)[j]
			}
		}
	}
	return false, 0, 0
}
