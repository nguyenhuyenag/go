package main

import "fmt"

/*
	- Chỉ có 1 loại duy nhât là for loop, nhưng có 3 cách viết khác nhau:

		Cách 1: Cú pháp giống C, Java, Python
			for i := 0; i < 5; i++ {
				fmt.Println(i)
			}

		Cách 2: Chỉ có điều kiện, không có phần khởi tạo và bước nhảy
			i := 0
			for i < 5 {
				fmt.Println(i)
				i++
			}

		Cách 3: Vòng lặp vô hạn (infinite loop)
			for {
				fmt.Println("Hello")
			}
*/

func loop1() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func loop2() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func loop3() {
	i := 0
	for {
		fmt.Println("Hello")
		i++
		if i == 5 {
		}
	}
}

/*
	- For each:

	- If you're looping over an slice, slice, string, or map, or reading from a channel,
	a range clause can manage the loop.
*/
func forEach() {
	arr := []int{6, 4, 3, 8, 5}
	for i, v := range arr {
		fmt.Println("Index=", i, ", Value=", v)
	}
}

func reverse() {
	arr := []int{1, 2, 3, 4, 5}
	// Reverse arr
	fmt.Println(arr)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println(arr)
}

func forArr() {
	arr := []int{10, 8, 2, 9, 0}
	for i := range arr {
		fmt.Printf("i=%d, a[%d]=%d\n", i, i, arr[i])
	}
}

func main() {
	// loop1()
	// loop2()
	// loop3()
	// forEach()
	// reverse()
	forArr()
}
