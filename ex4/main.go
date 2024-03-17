package main

import (
	"fmt"
	"os"
)

func main() {
	// This example reads the entire content of a file and writes it to another file.
	data, err := os.ReadFile("ex4/test.txt")

	if err != nil {
		panic(err)
	}
	fmt.Printf("File content: %s\n", data)

	if err := os.WriteFile("ex4/test_copy.txt", data, 0644); err != nil {
		panic(err)
	}

	fmt.Println("File copied")
}
