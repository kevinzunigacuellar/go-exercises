package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Read file using 100 byte buffer
	file, err := os.Open("files/test.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
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

}
