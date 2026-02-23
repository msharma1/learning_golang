package main

import (
	"bytes"
	"fmt"
)

func insertcomma(inputStr string) string {
	n := len(inputStr)
	var buffer bytes.Buffer
	if n <= 3 {
		return inputStr
	}
	for i := 0; i < n; i++ {
		if i != 0 && (i+1)%3 == 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(inputStr[i]))
	}
	return buffer.String()
}

func main() {
	fmt.Println(insertcomma("12345678900"))
}
