package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Now() : ", now)
	fmt.Println("Unix(): ", now.Unix())
	fmt.Println("UnixNano(): ", now.UnixNano())

	// Format
	fmt.Println("YYYY-MM-dd HH:mm:ss: ", now.Format("2006-01-02 15:04:05"))
	fmt.Println("dd/MM/YYYY: ", now.Format("02/01/2006"))
	fmt.Println("HH:ss: ", now.Format("15:04"))
	fmt.Println("ISO 8601: ", now.Format(time.RFC3339))
	//

}
