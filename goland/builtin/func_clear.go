package main

import "fmt"

/*
	- Hàm clear dùng để xóa dữ liệu trong map và slice.
	  + Với map, clear xóa tất cả các phần tử, khiến map trở thành map rỗng.
	  - Với slice, clear đặt tất cả các phần tử (từ đầu đến độ dài của slice)
		về giá trị zero của kiểu phần tử tương ứng.
*/
func clearSlide() {
	// Clear slice
	a := []int{1, 2, 3}
	fmt.Println(a)
	clear(a)
	fmt.Println(a)
}

func clearMap() {
	// Clear map
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(m)
	clear(m)
	fmt.Println(m)
}

func deleteKeyInMap() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println("Before:", m)
	delete(m, "a")
	fmt.Println("After:", m)
}

func clearAndEmptySlide()  {
	a1 := []int{1, 2, 3, 4}
	a2 := []int{1, 2, 3, 4}
	a3 := []int{1, 2, 3, 4}
	// Clear slice
	clear(a1)
	// Empty slice
	a2 = a2[:0]
	// nil slice
	a3 = nil
	fmt.Println("a1:", a1, len(a1))
	fmt.Println("a2:", a2, len(a2))
	fmt.Println("a3:", a3, len(a3))
}

func main() {
	// clearSlide()
	// clearMap()
	// deleteKeyInMap()
	clearAndEmptySlide()
}
