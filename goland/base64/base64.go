package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "Hello, World!"

	bytes := []byte(s)

	base64Encode := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println("Encode:", base64Encode)

	decodeString, err := base64.StdEncoding.DecodeString(base64Encode)
	// nil nghĩa là -> Không có giá trị / không trỏ tới dữ liệu nào.
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Decode:", string(decodeString))
}
