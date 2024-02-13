package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func getPath() string { 
	home, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }
	path := filepath.Join(home, ".simapp")

	return path
}

func getFiles(p string) []string {
	upgrades := []string{}
	files, err := os.ReadDir(p)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		log.Println(file.Name())
		upgrades = append(upgrades, file.Name())
	}
	return upgrades
}

func getLatestRelease() string {
    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://api.github.com/repos/elys-network/elys/releases", nil)
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Set("Accept", "application/vnd.github.v3+json")
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    var data interface{}
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Fatal(err)
    }
	releaseLast := data.([]interface{})[0].(map[string]interface{})["tag_name"].(string)
    return releaseLast
}

	func main() {
		log.Println("Automated Validator")

		db := getFiles(getPath())
		lr := getLatestRelease()

		for _, file := range db {
			if file == lr {
				log.Println("Latest Release Staged")
				return
			}
		}
		log.Println("Found new release :", lr)
		log.Println("Building latest binary....")
	}