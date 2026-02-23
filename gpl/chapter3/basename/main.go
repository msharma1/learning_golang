package main

import (
	"flag"
	"fmt"
)

func main() {
	inputPath := flag.String("path", "/home/abc", "Input path to find the basedir for")
	flag.Parse()
	path := *inputPath
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' && i != len(path)-1 {
			fmt.Println(string(path[i+1:]))
			return
		}
	}
}
