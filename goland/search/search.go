package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 4, 5, 7, 9}
	x := 6

	pos := sort.Search(len(arr), func(i int) bool {
		// tìm vị trí đầu tiên mà arr[i] >= x
		return arr[i] >= x
	})
	// pos = 4 (vì arr[4] = 7 >= 6)

	fmt.Printf("Chèn %d vào vị trí %d để giữ sắp xếp: %v\n", x, pos, arr)
}
