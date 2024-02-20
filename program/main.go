package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"time"
)

func main() {
	url := ReadUserInput("Enter the URL to generate QR code for (domain.com): ")
	//filePath := ReadUserInput("Enter the file path (leave blank to use default): ")
	//directory := ReadUserInput("Enter the directory to save the file (leave blank to use default): ")

	AvailableDomain := flag.Bool("check", false, "Check if the domain is available")
	flag.Parse() // Don't forget to parse the flags

	var finalPath string

	// Generate a default file name with current timestamp
	fileName := fmt.Sprintf("qr_code_%s_%s.png", url, time.Now().Format("20060102_150405"))
	home := homePath()

	if ReadUserInput("Enter the directory to save the file (leave blank to use default): ") == finalPath {
		// Use the user-specified directory
		finalPath = filepath.Join(home, finalPath, fileName)
	} else {
		// Use the default Windows Downloads Path
		defaultPath := GetWindowsDownloadsPath()
		finalPath = filepath.Join(defaultPath, fileName)
	}

	fmt.Println("Generating QR code for", url, "and saving to", finalPath)

	domain, err := DomainCheck(url) // Assuming this function checks the domain and returns a result and an error
	if err != nil {
		fmt.Println("Error checking domain:", err)
		return
	}

	if *AvailableDomain && domain != "" { // Assuming DomainCheck returns a non-empty string for valid domains
		fmt.Println("Domain is available:", domain)
	}

	b, err := GenerateQRCode(url) // Assuming this function generates the QR code and returns bytes and an error
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		return
	}

	err = Store(finalPath, b) // Assuming this function stores the QR code bytes to a file
	if err != nil {
		fmt.Println("Error saving QR code:", err)
		return
	}

	fmt.Println("QR code generated and saved successfully to", finalPath)
}
