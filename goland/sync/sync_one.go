package main

import (
	"fmt"
	"sync"
)

/*
	- Hàm chỉ chạy một lần duy nhất.
*/
func main() {
	var once sync.Once

	once.Do(func() {
		fmt.Println("Lệnh này sẽ được chạy")
	})

	once.Do(func() {
		fmt.Println("Lệnh này thì không được chạy")
	})
}
