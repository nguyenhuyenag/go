package main

import "fmt"

/*
	- Các phép toán bitwise:
		- AND:	&					bit = 1 khi cả 2 bit đều là 1
		- OR: 	|					bit = 1 khi ít nhất một bit là 1
		- XOR: 	^					bit = 1 khi 2 bit khác nhau
		- NOT: 	^ (đơn lẻ)
		- Left Shift:	<<
		- Right Shift: 	>>
 */
func main() {
	a := 6	//	110
	b := 3  // 	011
	fmt.Printf("%d & %d = %d\n", a, b, a & b) 	// = 010 = 2
	fmt.Printf("%d | %d = %d\n", a, b, a | b) 	// = 111 = 7
	fmt.Printf("%d ^ %d = %d\n", a, b, a ^ b)	// = 101 = 5
	fmt.Printf("NOT %d = %d\n", a, ^a)        	// = 001 = 1
	// fmt.Printf("%d << 1 = %d\n", a, a << 1)			// = 1100 = 12
	// fmt.Printf("%d >> 1 = %d\n", a, a >> 1)			// = 011 = 3
}
