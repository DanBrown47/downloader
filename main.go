package main

import (
	// Formatter
	"fmt"
	"io" //
	"log"
	"net/http" //http protocols

	//logger
	"os" // To create files
)

var err error

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage :<url> <filename>")
		os.Exit(1)
	}

	url := os.Args[1]      //Taking the url to download as input
	filename := os.Args[2] //Should split out the file name from url and add it here
	fmt.Println(url)

	err := Download(url, filename)
	if err != nil {
		log.Fatalf("Error downloading due to : %v", err)
	}

}

// Test http://speedtest.ftp.otenet.gr/files/test10Mb.db

// Downloader
func Download(url string, filename string) error {
	//create the file
	out, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating file might be due to : %v", err)
	}
	defer out.Close()

	//Download
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error while retriving the file form internet: %v", err)
	}
	defer resp.Body.Close()

	// Writing to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("Issue with saving the file might be due to : %v", err)
	}
	return nil
}
