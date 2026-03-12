package main

import (
	"fmt"
	"slices"
)

/*
	s := make([]T, length, capacity)

		Trong đó:
			- T → kiểu dữ liệu.
			- length → số phần tử hiện có trong slice.
			- capacity → dung lượng tối đa trước khi Go phải cấp phát lại bộ nhớ.

		*Lưu ý: Go dùng zero value nên nếu int mặc định sẽ là 0.

	> s = Clip(s) -> Giảm capacity của slide xuống bằng đúng len(s)

		s := make([]int, 3, 10)
		fmt.Println(len(s), cap(s)) // 3, 10
		s = slices.Clip(s)
		fmt.Println(len(s), cap(s)) // 3, 3
 */
func main() {
	s := make([]int, 5,	10)
	fmt.Println(s)
	// Insert by index
	s = slices.Insert(s, 1, 99)
	fmt.Println(s)
}
