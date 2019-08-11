package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	// numbers := []int{3, 2, 1}
	fmt.Println(numbers)
	fmt.Println(productAndDivision(&numbers))
}

func productAndDivision(numbers *[]int) *[]int {
	newArray := make([]int, len(*numbers))
	newArray[0] = 1

	for i := 1; i < len(*numbers); i++ {
		newArray[0] *= (*numbers)[i]
	}

	for i := 1; i < len(*numbers); i++ {
		newArray[i] = (newArray[i-1] / (*numbers)[i]) * (*numbers)[i-1]
	}

	return &newArray
}
