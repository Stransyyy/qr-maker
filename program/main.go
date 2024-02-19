package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	qr "github.com/skip2/go-qrcode"
)

func generate(url string) ([]byte, error) {
	return qr.Encode(url, qr.Medium, 256)
}

func store(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

func readUserInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return ""
}

func getWindowsDownloadsPath() string {
	user, err := user.Current()
	if err != nil {
		fmt.Printf("Unable to find the current user: %v\n", err)
		os.Exit(1)
	}
	downloadsPath := filepath.Join(user.HomeDir, "Downloads")

	if runtime.GOOS != "windows" && strings.Contains(strings.ToLower(os.Getenv("PATH")), "wsl") {
		// For WSL, attempt to construct the path assuming default Windows mount point
		downloadsPath = filepath.Join("/mnt", "c", "Users", user.Username, "Downloads")
	}
	return downloadsPath
}

func main() {
	// url := readUserInput("Enter the URL to generate QR code for: ")
	// filePath := readUserInput("Enter the file path (leave blank to use default): ")

	// if filePath == "" {
	// 	filePath = getWindowsDownloadsPath()
	// 	// Generate a default file name with current timestamp
	// 	fileName := fmt.Sprintf("qr_code_%s.png", time.Now().Format("20060102_150405"))
	// 	filePath = filepath.Join(filePath, fileName) // Correctly join the filePath and fileName
	// }

	// fmt.Println("Generating QR code for", url, "and saving to", filePath)

	// b, err := generate(url)
	// if err != nil {
	// 	fmt.Println("Error generating QR code:", err)
	// 	return
	// }

	// err = store(filePath, b)
	// if err != nil {
	// 	fmt.Println("Error saving QR code:", err)
	// 	return
	// }

	// fmt.Println("QR code generated and saved successfully to", filePath)

	domains, err := domain_check("www.VitalitySouth.com")
	if err != nil {
		fmt.Println("Error checking domain:", err)
		return
	}
	fmt.Println("Domains:", domains)
}

func URLshortened(url string) (string, error) {
	for url != "" {
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			return url, nil
		}
		url = "http://" + url
	}

	return url, nil
}

func ensureURLScheme(url string) (string, error) {
	if url == "" {
		return "", fmt.Errorf("empty URL")
	}
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	return url, nil
}
