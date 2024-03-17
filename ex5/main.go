package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	url := "https://example.com/"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalf("Failed to create a request: %s", err)
	}

	fmt.Println("Sending a request to:", url)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Failed to get a response: %s", err)
	}

	defer resp.Body.Close()
	f, err := os.Create("ex5/example.html")

	if err != nil {
		log.Fatalf("Failed to create a file: %s", err)
	}

	defer f.Close()

	fmt.Println("Downloading the file ...")

	if _, err := io.Copy(f, resp.Body); err != nil {
		log.Fatalf("Failed to copy the response to a file: %s", err)
	}
}
