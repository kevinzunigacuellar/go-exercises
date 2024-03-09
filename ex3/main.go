package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("ex3/test.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	// read file content
	buf := make([]byte, 100)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				// End of file, break the loop
				fmt.Print("\n")
				break
			} else {
				fmt.Println("Error reading file:", err)
				return
			}
		}

		// Print the read bytes as string
		fmt.Print(string(buf[:n]))
	}

	// print file content
	fmt.Println("Buffer size", len(buf))

}
