package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error occured reading the URL %v", err)
			os.Exit(1)
		}
		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error occured reading the content %v", err)
			os.Exit(1)
		}
		fmt.Printf("URL contents are: %s", content)
	}
}
