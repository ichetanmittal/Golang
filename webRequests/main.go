package main

import (
	"fmt"

	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Welcome to web requests in Go")
	const url = "https://google.com"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
	
	
}