package main

import (
	"io"
	"os"
)

func main() {
	// This example streams the content of a file to the standard output.
	file, err := os.Open("files/test.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	if _, err := io.Copy(os.Stdout, file); err != nil {
		panic(err)
	}

}
