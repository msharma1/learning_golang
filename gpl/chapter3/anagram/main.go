package main

import (
	"flag"
	"fmt"
	"reflect"
)

func anagram(input1 string, input2 string) bool {
	holder1 := make(map[rune]int, len(input1))
	holder2 := make(map[rune]int, len(input2))
	for _, alphabet := range input1 {
		holder1[alphabet]++
	}
	for _, alphabet := range input2 {
		holder2[alphabet]++
	}
	isAnagram := reflect.DeepEqual(holder1, holder2)
	return isAnagram
}

func main() {
	string1 := flag.String("s1", "abc", "first string")
	string2 := flag.String("s2", "cba", "first string")
	flag.Parse()
	if len(*string1) != len(*string2) {
		fmt.Println("unequal strings cannot be anagrams")
		return
	}
	fmt.Printf("Are %s and %s anagrams?: %t\n", *string1, *string2, anagram(*string1, *string2))
}
