package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var x uint64 = 0b10101010
	count := bits.OnesCount64(x)
	fmt.Printf("Số bit 1 trong %b là %d\n", x, count) // 1010101 có 4 số 1
	fmt.Println(bits.OnesCount64(0b11111111))         // 8
	fmt.Println(bits.OnesCount64(0b111111))           // 6
	fmt.Println(bits.OnesCount64(0b00))               // 0
}
