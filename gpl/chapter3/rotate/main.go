package main

import (
	"fmt"
)

/*
1,2,3,4,5,6

two parse
213456
216543
345612

345612
*/

func rotate(input []int, pos int) []int {
	result := make([]int, len(input))
	first := input[:pos]                  //first is [0,1]
	second := input[pos:]                 //second is [2,3,4,5]
	copy(result[:len(input)-pos], second) //input [:4] = [2,3,4,5] so input is [2,3,4,5,4,5]
	copy(result[len(input)-pos:], first)  //input[4:0] = [0,1] so input is [2,3,4,5,0,1]
	return result
}

func main() {
	input := []int{0, 1, 2, 3, 4, 5}
	result := rotate(input, 5)
	fmt.Println(result)
}
