package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	fmt.Println("Welcome to file handling in Go")
	content := "This needs to go in a file - LearnCodeOnline.in"

	
	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length is: ", length)

	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is \n", string(databyte))
}