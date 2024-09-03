package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Price int `json:"price"`
	Platform string `json:"platform"`
	Password string `json:"password"`
	Tags []string `json:"tags"`
}

func main() {
	
	lcoCourses := []User{
		{"Learn Code Online", 100, "lco.dev", "abc456", []string{"go", "js"}},
		{"MERN", 299, "lco.dev", "abc456", []string{"web dev", "react"}},
		{"Learn Code Online", 300, "lco.dev", "abc756", nil},
	}

	jsonFile, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonFile))
}

func DecodeJson() {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}
	


}