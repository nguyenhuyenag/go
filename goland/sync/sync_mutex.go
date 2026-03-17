package main

import (
	"fmt"
	"sync"
)

var count int

var mu sync.Mutex

func incr() {
	mu.Lock()
	count++
	mu.Unlock()
}

/*
	- Mutex cho phép chỉ một goroutine truy cập vùng code tại một thời điểm.
 */
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			incr()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(count)
}