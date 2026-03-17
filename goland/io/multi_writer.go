package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file1, err := os.Create("test_output1.txt")
	if err != nil {
		panic(err)
	}
	defer file1.Close()

	file2, err := os.Create("test_output2.txt")
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	mw := io.MultiWriter(file1, file2, os.Stdout)
	fmt.Fprintln(mw, "Dòng này vừa ghi ra màn hình, vừa ghi vào file.")
}
