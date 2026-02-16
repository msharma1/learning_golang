package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error occured reading the URL %v", err)
			os.Exit(1)
		}
		content, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error occured reading the content %v", err)
			os.Exit(1)
		}
		fmt.Printf("URL contents are: %s", content)
		fmt.Printf("HTTP status code: ", response.Status)
	}
}
