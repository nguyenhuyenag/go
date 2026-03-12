package main

import (
	"fmt"
	"math"
)

func main() {
	// (real + imag*i)
	c := complex(3, 4)
	fmt.Println(c)
	r := math.Sqrt(real(c)*real(c) + imag(c)*imag(c))
	fmt.Println(r)
}
