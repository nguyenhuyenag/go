package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	- Đọc dữ liệu từ bàn phím.
*/
func readLine() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	text, _ := reader.ReadString('\n') // <- Đọc đến khi gặp ký tự xuống dòng
	fmt.Println("Output: ", text)
}

func readLines() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter lines (-1 to exit):")
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		if line == "-1" {
			break
		}
		fmt.Printf("Line %d: %s\n", i, line)
		i++
	}
}

func main() {
	// readLine()
	readLines()
}
