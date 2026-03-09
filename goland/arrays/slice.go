package main

import "fmt"

/*
	s := make([]T, length, capacity)

	Trong đó:
		- T → kiểu dữ liệu.
		- length → số phần tử hiện có trong slice.
		- capacity → dung lượng tối đa trước khi Go phải cấp phát lại bộ nhớ.

	*Lưu ý: Go dùng zero value nên nếu int mặc định sẽ là 0.
 */
func main() {
	s := make([]int, 5,	10)
	fmt.Println(s)
}
