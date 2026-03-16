package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// readAll()
	readLineByLine()
}

func readLineByLine() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println("Project directory:", dir)

	filePath := filepath.Join(dir, "data", "currency.csv")
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Đọc header trước
	header, err := reader.Read()
	if err != nil {
		panic(err)
	}
	fmt.Println("Header:", header)

	// Đọc từng dòng
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Printf("(%s=%s, %s=%s, %s=%s)\n",
			header[0], row[0],
			header[1], row[1],
			header[2], row[2],
		)
	}
}

func readAll() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println("Project dir:", dir)

	filePath := filepath.Join(dir, "data", "currency.csv")
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, row := range records {
		fmt.Println(row[0], row[1], row[2])
	}
}
