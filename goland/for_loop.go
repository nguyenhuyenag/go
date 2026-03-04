package main

import "fmt"

/*
	- Chỉ có 1 loại duy nhât là for loop, nhưng có 3 cách viết khác nhau:

	1. Cách 1: Cú pháp giống C, Java, Python
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}

	2. Cách 2: Chỉ có điều kiện, không có phần khởi tạo và bước nhảy
		i := 0
		for i < 5 {
			fmt.Println(i)
			i++
		}

	3. Cách 3: Vòng lặp vô hạn (infinite loop)
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

func main() {
	loop1()
	// loop2()
	// loop3()
}
