package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Intn(max - min) + min
	fmt.Println(rand.Intn(100))	//	0 → 99
	fmt.Println(rand.Float64()) 	// 	0.0 → 1.0
}
