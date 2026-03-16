package main

import "fmt"

func main() {
	s := "abcd"
	r := []rune(s) 				// Convert to unicode array: [97, 98, 99, 100]
	chars := []byte(s)
	fmt.Println(r)
	fmt.Println(chars)
	fmt.Println(string(r))
	fmt.Println(string(chars))
}
