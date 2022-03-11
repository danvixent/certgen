package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
)

// CERTIFICATES AUXILIARY FUNCTIONS

// efficiently download a file from url
func downloadFile(url string, filepath string) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		log.WithError(err).Fatal("failed to create download file")
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.WithError(err).Fatal("failed to send file download request")
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.WithError(err).Fatal("failed to write downloaded file to out")
	}
}

func getAppData() string {
	dir := ""
	switch {
	case runtime.GOOS == "windows":
		dir = os.Getenv("LocalAppData")
		//return filepath.Join(dir, "sserve") + "\\"
	case os.Getenv("XDG_DATA_HOME") != "":
		dir = os.Getenv("XDG_DATA_HOME")
	case runtime.GOOS == "darwin":
		dir = os.Getenv("HOME")
		if dir == "" {
			return ""
		}
		dir = filepath.Join(dir, "Library", "Application Support")
	default: // Linux/Unix
		dir = os.Getenv("HOME")
		if dir == "" {
			return ""
		}
		dir = filepath.Join(dir, ".local", "share")
	}
	appData := filepath.Join(dir, "certgen")
	err := os.MkdirAll(appData, os.ModePerm)
	if err != nil {
		log.WithError(err).Fatal("failed to create appdata directory")
	}
	return appData + "/"
}

// mkcert to generates certificates
func mkcert() {
	// set the right executable according to the system
	exeURL := "https://github.com/FiloSottile/mkcert/releases/download/v1.2.0/"
	file := ""
	switch runtime.GOOS {
	case "darwin":
		file = "mkcert-v1.2.0-darwin-amd64"
	case "linux":
		file = "mkcert-v1.2.0-linux-amd64"
	case "windows":
		file = "mkcert-v1.2.0-windows-amd64.exe"
	default:
		log.Fatal("Your system is not supported. Sorry.")
		os.Exit(1)
	}

	// download mkcert binaries
	appData := getAppData()
	downloadFile(exeURL+file, appData+file)

	// make binary executable
	err := os.Chmod(appData+file, 0755)
	if err != nil {
		log.WithError(err).Fatal("failed to chmod appdate file to 0755")
	}

	// generate the certificate
	if _, err := exec.Command(appData+file, "-install", "-cert-file", appData+"localhost.crt",
		"-key-file", appData+"localhost.key", "localhost").Output(); err != nil {
		log.WithError(err).Fatal("failed to execute mkcert")
	}

	log.Println("Generated certificates in " + appData + ".")
}

// check if file exists
func exist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		os.Exit(1)
		return false
	}
}

func getCert() (string, string) {
	appData := getAppData()
	// ensure that the certificate files exists
	if !exist(appData+"localhost.crt") || !exist(appData+"localhost.key") {
		mkcert()
	} else {
		log.Println("Using certificates in " + appData + ".")
	}
	return appData + "localhost.crt", appData + "localhost.key"
}

// CLI interface
func main() {
	crt, key := getCert()
	err := os.Setenv("LOCALHOST_CRT_FILE", crt)
	if err != nil {
		log.WithError(err).Fatalf("failed to set LOCALHOST_CRT_FILE env var")
	}

	err = os.Setenv("LOCALHOST_KEY_FILE", key)
	if err != nil {
		log.WithError(err).Fatalf("failed to set LOCALHOST_KEY_FILE env var")
	}

	fmt.Println("LOCALHOST_CRT_FILE", crt)
	fmt.Println("LOCALHOST_KEY_FILE", key)
}
