package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// Object to JSON
func convertToJSON() {
	p := Person{"Alice", 25}
	// data, _ := json.Marshal(p)
	data, _ := json.MarshalIndent(p, "", "  ")
	// fmt.Println(reflect.TypeOf(data))
	jsonData := string(data)
	fmt.Println(jsonData)
}

func JSONToObject() {
	str := `{"Name":"Alice","Age":25}`
	var p Person
	err := json.Unmarshal([]byte(str), &p)
	if err != nil {
		return
	}
	fmt.Println(p.Name, p.Age)
}

/*
	Hàm							Chức năng
	json.Marshal()				Struct → JSON
	json.Unmarshal()			JSON → struct
	json.MarshalIndent()		JSON đẹp
	json.NewEncoder()			Ghi JSON vào stream
	json.NewDecoder()			Đọc JSON từ stream
*/
func main() {
	convertToJSON()
	// JSONToObject()
}
