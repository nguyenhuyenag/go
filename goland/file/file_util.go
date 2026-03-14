package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Current directory
func currentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println("Current directory:", dir)
	return dir
}

// Lấy thư mục của file đang chạy
func fileExecutable() {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	dir := filepath.Dir(exe)
	fmt.Println(dir)
}

// Lấy thư mục của source file
func sourceDir() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)
}

func main() {
	// currentDir()
	// fileExecutable()
	// sourceDir()
}
