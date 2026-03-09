package main

import "fmt"

/*
	%v   	giá trị mặc định
	%T   	kiểu dữ liệu
	%d   	integer
	%f   	float
	%.2f 	float 2 số lẻ
	%s   	string
	%t   	boolean
	%p   	pointer
	%q   	string có dấu ""
 */
func main() {
	x := 65
	fmt.Printf("%d\n", x) // 65
	fmt.Printf("%b\n", x) // 1000001
	fmt.Printf("%c\n", x) // A
}
