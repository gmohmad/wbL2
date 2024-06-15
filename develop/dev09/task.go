package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
=== wget utility ===

Implement the wget utility with the ability to download entire sites

The program must pass all tests. The code must pass go vet and golint checks.
*/

func main() {
	// if no arguments were provided print this message and return
	if len(os.Args) < 2 {
		fmt.Println("Usage: wget <URL>")
		return
	}

	url := os.Args[1]
	err := downloadFile("downloaded_page.html", url)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Downloaded:", url)
	}
}

func downloadFile(filename string, url string) error {
	// Create the file
	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error downloading file: %v", err)
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("error copying data to file: %v", err)
	}

	return nil
}
