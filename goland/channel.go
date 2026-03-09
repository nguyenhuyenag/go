package main

import "fmt"

/*
	- Channel là cơ chế trong Go dùng để gửi và nhận dữ liệu giữa các goroutine một cách an toàn.
	- Nó là công cụ chính của concurrency trong Go.
	- Đừng chia sẻ memory rồi lock, mà hãy truyền dữ liệu qua channel.
*/
func main() {
	ch := make(chan int)

	go func() {
		ch <- 10
	}()

	v := <-ch

	fmt.Println(v)
}
