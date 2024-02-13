package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"io"
)

func getPath() string { 
	home, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }
	path := filepath.Join(home, ".simapp")

	return path
}

func getFiles(p string) {
	files, err := os.ReadDir(p)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		log.Println(file.Name())
	}

}

func getLatestRelease() {
	resp, err := http.Get("https://github.com/elys-network/elys/releases")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(body)
	
}

	func main() {
		log.Println("Automated Validator")
		getFiles(getPath())
		getLatestRelease()
	}