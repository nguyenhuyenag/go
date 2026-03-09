package main

import (
	"fmt"
)

/*
	> Cách 1:

		if err := file.Chmod(0664); err != nil {
			log.Print(err)
			return err
		}

	> Cách 2:

		f, err := os.Open(name)
		if err != nil {
			return err
		}
 */

func main() {
	a := 10
	if a % 2 == 0 {
		fmt.Println("a is even")
	} else {
		fmt.Println("a is odd")
	}
}
