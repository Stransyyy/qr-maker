package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"time"
)

func main() {
	url := ReadUserInput("Enter the URL to generate QR code for (domain.com): ")
	filePath := ReadUserInput("Enter the file path (leave blank to use default): ")

	AvailableDomain := flag.Bool("check", false, "Check if the domain is available")

	if filePath == "" {
		filePath = GetWindowsDownloadsPath()
		// Generate a default file name with current timestamp
		fileName := fmt.Sprintf("qr_code_%s_%s.png", url, time.Now().Format("20060102_150405"))
		filePath = filepath.Join(filePath, fileName) // Correctly join the filePath and fileName
	}

	fmt.Println("Generating QR code for", url, "and saving to", filePath)

	domain, err := DomainCheck(url)
	if err != nil {
		fmt.Println("Error checking domain:", err)
		return
	}

	if *AvailableDomain {
		for domains := range domain {
			fmt.Println("Domain is  available:", domains)
		}
	}

	b, err := Generate(domain)
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		return
	}

	err = Store(filePath, b)
	if err != nil {
		fmt.Println("Error saving QR code:", err)
		return
	}

	fmt.Println("QR code generated and saved successfully to", filePath)

}
