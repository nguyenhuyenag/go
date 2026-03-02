package main

import "fmt"

/*
	Kiểu    Ý nghĩa
	int     Số nguyên (phổ biến nhất)
	int8    -128 → 127 			(kiểu byte)
	int16   Số nguyên 16 bit
	int32   Số nguyên 32 bit 	(kiểu rune)
	int64   Số nguyên 64 bit
	uint    Số nguyên không âm
*/
func intType() {
	var a int = 10
	var b int8 = 11
	var c uint = 12
	d := 13 		// Khai báo biến mới + gán giá trị. Nếu định nghĩa lần nữa `d := ...` sẽ lỗi
	fmt.Println(a, b, c, d)
}

func floatType() {
	var a float32 = 3.14
	var b float64 = 2.71828
	c := 1.5 // Go mặc định là float64 nếu có dấu thập phân

	fmt.Println(a, b, c)

	v1 := 1.01
	v2 := 0.99
	fmt.Println(v1 - v2)
}

func main() {
	// intType()
	floatType()
}
