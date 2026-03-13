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

	> Concat(): Nối nhiều slice lại thành một slice mới

		a := []int{1, 2}
		b := []int{3, 4}

		r := slices.Concat(a, b)

	> Contains(slice, v): Kiểm tra xem slice có chứa phần tử v hay không.
	> ContainsFunc(): Kiểm tra xem slice có phần tử nào thỏa điều kiện hay không.

		s := []int{1, 3, 5, 7}
		found := slices.ContainsFunc(s, func(x int) bool {
			return x % 2 == 0 // Kiểm tra xem có số chẵn nào hay không?
		})

	> Delete(s, start, end): Xóa một đoạn phần tử [start, end] trong slice theo index.
	> Delete(s, i, i + 1): Xóa 1 phần tử ở vị trí i.
	> DeleteFunc(): Xóa các phần tử thỏa điều kiện.

		s := []int{1, 2, 3, 4, 5}
		s = slices.DeleteFunc(s, func(x int) bool {
			return x % 2 == 0 // Xóa các số chẵn
		})

	> Equal(): Kiểm tra 2 slides có bằng nhau hay không?
	> EqualFunc():

		s1 := []string{"Go", "Java"}
		s2 := []string{"go", "JAVA"}

		// So sánh không phân biệt hoa thường
		equal := slices.EqualFunc(s1, s2, func(x, y string) bool {
			return strings.EqualFold(x, y)
		})

		// So sánh các User theo ID
		equal := slices.EqualFunc(s1, s2, func(x, y User) bool {
			return x.ID == y.ID
		})

	> Index(slice, v): Tìm index của v.
	> IndexFunc(): Tìm index của phần tử thỏa điều kiện.

		s := []int{1, 3, 5, 7}
		index := slices.IndexFunc(s, func(x int) bool {
			return x % 2 == 0 // Tìm index của số chẵn đầu tiên
		})

	> Insert(s, index, value): Insert by index.

	> Repeat(s, n): Lặp lại slice n lần.

	> Replace(s, i, j, v1, v2,...): Xóa đoạn [i, j) và thay bằng v1,v2,..

	> Reverse(s): Đảo mảng

		s[i], s[j] = s[j], s[i]

	> s = Clip(s) -> Giảm capacity của slice xuống bằng đúng len(s)

		s := make([]int, 3, 10)
		fmt.Println(len(s), cap(s)) // 3, 10
		s = slices.Clip(s)
		fmt.Println(len(s), cap(s)) // 3, 3
*/
func main() {
	//s := []int{4, 6, 2, 0, 34, 2, 9}
	//fmt.Println("Contains:", slices.Index(s, 2))
	// fmt.Println(s)
	// Insert by index
	//s = slices.Insert(s, 1, 99)
	// fmt.Println(s)
	a := []int{1, 2, 3}
	slices.Reverse(a)
	fmt.Println(a)
	//
	//b := slices.Repeat(a, 3)
	//
	//fmt.Println(b)
	//words := []string{"Go", "Lang"}
	//result := slices.Repeat(words, 2)
	//fmt.Println(result)
	//a := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	//b := slices.Replace(a, 1, 4, -1, -2, -3, -4)
	//b1 := slices.Replace(a, 1, len(b), -1)
	//
	//fmt.Println(b)
	//fmt.Println(b1)
}
