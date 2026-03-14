package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println("Current directory:", dir)

	path := filepath.Join(dir, "file", "data.csv")
	file, err := os.Open(path)

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
		// fmt.Println(len(row))
		// fmt.Println(reflect.TypeOf(row), row)
		fmt.Println(row[0], row[1], row[2])
	}

}
