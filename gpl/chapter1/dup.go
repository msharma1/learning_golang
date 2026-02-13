// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("Give at least one file name next time. Exiting now...")
	} else {
		for _, filename := range files {
			counter := make(map[string]int)
			currentFile, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			counter = countLines(currentFile, counter)
			currentFile.Close()
			fmt.Printf("In file: %s:\n", filename)
			for lineStr, numOfOccurance := range counter {
				fmt.Printf("String %s occured %d times in above file.\n", lineStr, numOfOccurance)
			}
		}
	}

}

func countLines(currentFile *os.File, counter map[string]int) map[string]int {
	input := bufio.NewScanner(currentFile)
	for input.Scan() {
		counter[input.Text()]++
	}

	if input.Err() != nil {
		fmt.Println("Error in reading the file. Will continue to the next file")
	}
	return counter
}
