package main

import "fmt"

func insertcomma(inputStr string) string {
	n := len(inputStr)
	holder := inputStr
	if n <= 3 {
		return inputStr
	}
	for i := n - 1; i >= 0; i-- {
		if (i+1)%3 == 0 {
			holder = holder[:i] + "," + holder[i:]
		}
	}
	return holder
}

func main() {
	fmt.Println(insertcomma("12345678900"))
}
