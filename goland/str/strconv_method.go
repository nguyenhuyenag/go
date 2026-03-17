package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// String → Number
func stringToNumber() {
	s := "123"
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(n) // 123
}

// String → int (custom base)
func stringToInt() {
	n, _ := strconv.ParseInt("1010", 2, 64)
	fmt.Println(n) // 10
}

// String → float
func stringToFloat() {
	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println(f)
}

// String → bool
func stringToBool() {
	b, _ := strconv.ParseBool("true")
	fmt.Println(b)
}

// Number → String
func numberToString() {
	n := 123             // int
	s := strconv.Itoa(n) // string
	fmt.Println("Type(n):", reflect.TypeOf(n))
	fmt.Println("Type(s):", reflect.TypeOf(s))
	fmt.Println(s)
}

// int → string (custom base)
func intToString() {
	var n int64 = 10
	s := strconv.FormatInt(n, 2) // Binary
	fmt.Println(s)                // "1010"
}

func floatToString() {
	/*
		'f'		format
		2		số chữ số sau dấu phẩy
		64		bit
	 */
	s := strconv.FormatFloat(3.14564, 'f', 2, 64)
	fmt.Println(s)
}

func quoteUnquote() {
	// Biến string literal có dấu nháy kép và escape ký tự đặc biệt
	s1 := strconv.Quote("hello\nworld")
	fmt.Println(s1)

	// Chuyển ngược lại: String có dấu nháy + escape → string bình thường
	s2, _ := strconv.Unquote("\"hello\\nworld\"")
	fmt.Println(s2)
}

func main() {
	// stringToNumber()
	// stringToInt()
	// stringToFloat()
	// stringToBool()
	// numberToString()
	// intToString()
	// floatToString()
	quoteUnquote()
}
