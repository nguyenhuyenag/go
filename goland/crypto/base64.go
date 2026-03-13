package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "Hello, World!"

	bytes := []byte(s)

	encodeStr := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println("Encode:", encodeStr)

	decodeBytes, err := base64.StdEncoding.DecodeString(encodeStr)
	// nil nghĩa là -> Không có giá trị / không trỏ tới dữ liệu nào.
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Decode:", string(decodeBytes))
}
