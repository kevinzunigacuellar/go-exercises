package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("ex4/test.txt")

	if err != nil {
		panic(err)
	}

	fmt.Printf("File content: %s\n", data)

	if err := os.WriteFile("ex4/copy.txt", data, 0644); err != nil {
		panic(err)
	}

	fmt.Println("File copied")
}
