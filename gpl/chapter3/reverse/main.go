package main

import "fmt"

func reverseWithSlice(input []int) []int {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func reverseWithArrayPointer(input *[4]int) {
	for i, j := 0, len(*input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}

func main() {
	input := [...]int{1, 2, 3, 4}
	reverseWithSlice(input[:])
	fmt.Println(input)
	reverseWithArrayPointer(&input)
	fmt.Println(input)
}
