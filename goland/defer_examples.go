package main

import (
	"fmt"
	"os"
)

/*
	- Defer: Trì hoãn việc gọi một hàm cho đến khi hàm hiện tại kết thúc.

	-> Defer = chạy sau cùng, trước khi hàm return.

	- Defer lưu function, không lưu argument
*/

func deferBasic() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

/*
	- Thứ tự chạy (LIFO).

	-> Nếu có nhiều defer, chúng chạy ngược thứ tự khai báo
*/
func lifo() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println("Start")
}

/*
	- Dùng nhiều nhất: Đóng resource.
*/
func closeResource() {
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}
	defer file.Close()
}

func anonymous() {
	x := 10

	defer func() {
		fmt.Println(x)
	}()

	x = 20
}

func main() {
	// deferBasic()
	// lifo()
	// closeResource()
	anonymous()
}
