package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/get", handleGet)

	fmt.Println("Starting server at http://localhost:8000")
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request from %s", r.RemoteAddr)
	fmt.Fprintf(w, "Hello from the server!")
}

// PerformGetRequest can be called to test the server
func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		log.Fatalf("Error making request: %s\n", err)
	}
	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %s\n", err)
	}

	fmt.Println("Content:", string(content))
}