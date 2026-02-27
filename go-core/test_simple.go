package main

import "fmt"

func main() {

	a := [3]int{1, 2, 3}

	for i, v := range a {
		fmt.Printf("Phan tu thu %d la: %d\n", i, v)
	}

}
