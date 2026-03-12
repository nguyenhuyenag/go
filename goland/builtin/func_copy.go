package main

import "fmt"

/*
	copy(dst, src): -> Sao chép dữ liệu từ src -> dst.
		return min(len(src),len(dst)) // Số lượng phần tử đã được sao chép.
 */
func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 6)
	n := copy(dst, src)
	fmt.Println(dst) // [1 2 3 4 5 0]
	fmt.Println(n)
}
