package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

func Store(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

func ReadUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text(), fmt.Errorf("error reading input: %v", scanner.Err())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return "", nil
}

// Presents the user with directory options and allows selection or custom path entry.
func chooseDownloadDirectory() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	homeDir := usr.HomeDir
	if isWSL() {
		fmt.Println("Detected WSL. Adjusting for Windows file system.")
		homeDir = fmt.Sprintf("/mnt/c/Users/%s", usr.Username) // Adjust for WSL
	}

	// Define potential download directories here, adjust as necessary.
	directoryOptions := map[string]string{
		"1": filepath.Join(homeDir, "Downloads"),
		"2": filepath.Join(homeDir, "Documents"),
		"3": "Enter a custom directory path",
	}

	fmt.Println("Where would you like to download the file?")
	for key, option := range directoryOptions {
		if key == "3" {
			fmt.Printf("%s. %s\n", key, option)
		} else {
			fmt.Printf("%s. %s\n", key, filepath.Base(option))
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Select an option or enter a custom path: ")
	if scanner.Scan() {
		choice := scanner.Text()

		// Handle custom directory input
		if choice == "3" || !strings.Contains("123", choice) {
			fmt.Print("Enter the full path to the directory: ")
			if scanner.Scan() {
				customPath := scanner.Text()
				return customPath, nil
			}
		} else {
			// Use selected predefined option
			selectedPath, exists := directoryOptions[choice]
			if exists {
				return selectedPath, nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("failed to select a directory")
}

func defaultDownloadPath() (string, error) {
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
	return downloadsPath, nil
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

func isWSL() bool {
	if runtime.GOOS != "linux" {
		return false
	}

	versionContents, err := os.ReadFile("/proc/version")
	if err != nil {
		return false
	}

	versionInfo := strings.ToLower(string(versionContents))
	return strings.Contains(versionInfo, "microsoft") || strings.Contains(versionInfo, "microsoft-standard")
}
