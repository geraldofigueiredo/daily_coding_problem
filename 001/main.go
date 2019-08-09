package main

import "fmt"

func main() {
	numbers := []int32{1, 4, 5, 6, 9, 13, 36, 78, 45}
	searchAddup(&numbers, 20)
}

func searchAddup(numbers *[]int32, k int32) {
	fmt.Println(*numbers)
}
