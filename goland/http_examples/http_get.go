package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var url = "https://microsoftedge.github.io/Demos/json-dummy-data/64KB.json"

type Person struct {
	Name     string  `json:"name"`
	Language string  `json:"language"`
	ID       string  `json:"id"`
	Bio      string  `json:"bio"`
	Version  float64 `json:"version"`
}

func main() {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var data []Person

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Total:", len(data))
	fmt.Println(data[0].Name)
}
