package main

import (
	"encoding/json"
	"fmt"
	"reflect"
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
	fmt.Println(reflect.TypeOf(data))
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

func main() {
	convertToJSON()
	// JSONToObject()
}
