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

func ReadUserInput(prompt string) string {
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

func userDownloadsPath(directory string) string {
	user, err := user.Current()
	if err != nil {
		fmt.Printf("Unable to find the current user: %v\n", err)
		os.Exit(1)
	}
	return filepath.Join(user.HomeDir, directory)
}

func homePath() string {
	user, err := user.Current()
	if err != nil {
		fmt.Printf("Unable to find the current user: %v\n", err)
		os.Exit(1)
	}
	return user.HomeDir

}

func GetWindowsDownloadsPath() string {
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

func URLshortened(url string) (string, error) {
	for url != "" {
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			return url, nil
		}
		url = "http://" + url
	}

	return url, nil
}
