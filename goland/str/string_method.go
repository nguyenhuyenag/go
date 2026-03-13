package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "    Hello, World!               "
	s2 := "!!!!!Hello World!!"

	fmt.Println("s1:", s1)

	// Loại bỏ khoảng trắng 2 đầu
	fmt.Println("TrimSpace:", strings.TrimSpace(s1))

	// Loại bỏ khoảng trắng bên trái
	fmt.Println("TrimLeft:", strings.TrimLeft(s1, " "))

	// Loại bỏ khoảng trắng bên phải
	fmt.Println("TrimRight:", strings.TrimRight(s1, " "))

	// Loại bỏ ký tự được chỉ định ở 2 đầu
	fmt.Println("Trim '!':", strings.Trim(s2, "!"))

}
