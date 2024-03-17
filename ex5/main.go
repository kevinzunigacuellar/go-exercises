package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://example.com/", nil)

	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	flags := os.O_CREATE | os.O_WRONLY
	f, err := os.OpenFile("ex5/response.txt", flags, 0644)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err := io.Copy(f, resp.Body); err != nil {
		panic(err)
	}
}
