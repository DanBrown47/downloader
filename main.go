package main

import (
	// Formatter
	"fmt"
	"io"
	"net/http"
	"os" // To create files
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage :<url> <filename>")
		os.Exit(1)
	}

	url := os.Args[1] //Taking the url to download as input
	filename := os.Args[2]
	fmt.Println(url)

	err := Download(url, filename)

	if err != nil {
		panic(err)
	}

}

// Test http://speedtest.ftp.otenet.gr/files/test10Mb.db

// Downloader
func Download(url string, filename string) error {
	//create the file
	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	//Download
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Writing to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
	return nil
}
