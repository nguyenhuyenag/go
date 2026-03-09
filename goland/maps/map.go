package main

import "fmt"

func main() {

	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}

	// Show all
	fmt.Println("> All:", timeZone)

	// ???
	fmt.Println("> One:", timeZone["EST"])

	/*
		value, ok := timeZone["EST"]

		Trong đó:
			- value → giá trị của key
			- ok → true nếu key tồn tại
			- ok → false nếu key không tồn tại
	*/
	if value, ok := timeZone["EST"]; ok {
		// Biến value và ok chỉ tồn tại trong if block.
		fmt.Println("Value:", value)
	}

	// Delete
	delete(timeZone, "EST")

	// Loop
	fmt.Println("> Loop:")
	for k, v := range timeZone {
		fmt.Printf("Key=%s, Value=%d\n", k, v)
	}

}
