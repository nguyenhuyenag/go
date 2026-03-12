package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// string <=> int
func main() {
	var n1 = 123
	var s = strconv.Itoa(n1)
	n2, _ := strconv.Atoi(s)
	fmt.Println(reflect.TypeOf(n2), ":", n1)
	fmt.Println(reflect.TypeOf(s), ":", n2)
}
