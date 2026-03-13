package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte("hello")
	hash := md5.Sum(data)
	fmt.Printf("MD5: %x\n", hash)
}
