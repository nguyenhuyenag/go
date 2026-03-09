package main

import "fmt"

/*
Một số format phổ biến:

	Ký hiệu		Ý nghĩa
	%d			int
	%f			float
	%s			string
	%v			mọi kiểu
	%.2f		float 2 chữ số thập phân
*/
func printType() {
	// Không xuống dòng
	fmt.Println("(1) Không xuống dòng")
	fmt.Println("World")
	fmt.Print("Hello ")
	fmt.Print("World")

	// Tự động xuống dòng
	fmt.Println("(2) Tự động xuống dòng")
	fmt.Println("Hello")
	fmt.Println("World")

	// Format
	fmt.Printf("Tên: %s, Tuổi: %d\n", "An", 20)

	// Trả về chuỗi, không in ra màn hình
	s1 := fmt.Sprint("Hello", 10)
	fmt.Println(s1)

	// Format rồi trả về chuỗi
	s2 := fmt.Sprintf("Tuổi: %d", 20)
	fmt.Println(s2)

	// Đội dài chuỗi
	fmt.Println(len("Hello World"))

	// String thực chất là một mảng byte (UTF-8)
	fmt.Println("MMHello World"[1]) //  = 'e' nhưng Go không in 'e' mà in ra ascii('e') = 101
	fmt.Println(string("Hello World"[1]))

	fmt.Println("Hello " + "World")
}

func stringType() {
	var s1 string = "Hello"
	var s2 string = "World"
	s3 := "Go is great!"

	println(s1, s2, s3)

	// Kết hợp chuỗi
	full := s1 + " " + s2 + "!"
	println(full)

	// Chuỗi đa dòng
	multiline := `This is a
						multiline string.`
	println(multiline)

}

func main() {
	printType()
	// stringType()
}
