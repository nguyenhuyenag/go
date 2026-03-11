package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var path = "E:/Dev/Projects/Github/go/goland/io/file_utils.go"

func readFileByLine() {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		fmt.Printf("Line %d: %s\n", i, scanner.Text())
	}
}

func readToByteArray() {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Byte[]: ", data) // []byte
	fmt.Println(string(data))     // chuyển sang string
}

// Đọc bằng io.ReadAll
func readToByte2() {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)

	fmt.Println(data)
}

func validate() {
	_, err := os.Stat(path)
	// _, err := os.Stat("data.txt")

	if err == nil {
		fmt.Println("File exists")
	} else if os.IsNotExist(err) {
		fmt.Println("File does not exist")
	} else {
		fmt.Println("Other error:", err)
	}
}

func main() {
	validate()
	// readFileByLine()
	// readToByteArray()
	// readToByte2()
}
