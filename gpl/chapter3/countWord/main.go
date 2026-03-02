package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	counts := make(map[string]int)
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		counts[word]++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading input: %v\n", err)
	}

	fmt.Printf("Word\t\tFrequency\n")
	fmt.Println("----\t\t---------")
	for word, n := range counts {
		fmt.Printf("%-15s %d\n", word, n)
	}

}
