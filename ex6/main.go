package main

import (
	"io"
	"os"
)

func main() {

	file, err := os.Open("ex6/model.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(os.Stdout, file); err != nil {
		panic(err)
	}

}
