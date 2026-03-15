package main

import (
	"io"
	"os"
)

func main() {
	src, _ := os.Open("a.txt")
	dst, _ := os.Create("b.txt")

	_, err := io.Copy(dst, src)
	if err != nil {
		return
	}
}

