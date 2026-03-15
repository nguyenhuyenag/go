package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Join([]string{"a", "b", "c"}, " "))
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(strings.Join([]string{"a", "b", "c"}, " "))
}
