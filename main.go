package main

import (
	// Formatter
	"fmt"
	"io" //
	"log"
	"net/http" //http protocols
	"strings"

	//logger
	"os" // To create files
)

var err error

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage :<options> <url> ")
		os.Exit(1)
	}

	url := os.Args[1] //Taking the url to download
	urlsplit := strings.Split(url, string('/'))
	filename := urlsplit[len(urlsplit)-1] //Split out the file name from url
	fmt.Println(filename)

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
