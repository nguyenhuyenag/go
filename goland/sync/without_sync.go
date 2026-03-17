package main

import "fmt"

var counter int

func increment() {
	counter++
}

/*
	- Khi nhiều goroutine cùng sửa một biến, kết quả có thể sai.
*/
func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}

	fmt.Println(counter) // = 997 -> Kết quả thường không phải 1000 vì có race condition.
}
