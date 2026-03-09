package main

import "fmt"

// Slice -> Tự động tăng kích thước
func baseList()  {
	list := []string{"a", "b", "c"}

	// Hiển thị giá trị tại vị trí cụ thể
	fmt.Println(list[0])

	// Hiển thị toàn bộ list
	fmt.Println(list)

	// Thay đổi giá trị
	list[0] = "d"
	fmt.Println(list)

	// Thêm vị trí mới
	list = append(list, "e")
	fmt.Println(list)

	// Đội dài của list
	println(len(list))

	// Returns the capacity
	println(cap(list))

	//
	fmt.Println("s[1:3] =", list[1:3])

	// Duyệt list
	for i, v := range list {
		fmt.Println(i, v)
	}

}

func main() {
	baseList()
}
