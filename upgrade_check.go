package main

import (
	"log"
	"os"
	"path/filepath"
)

func getStagedBinaries() []string {
    home, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }
    path := filepath.Join(home, ".simapp")

    upgrades := []string{}
    files, err := os.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }
    for _, file := range files {
        log.Println(file.Name())
        upgrades = append(upgrades, file.Name())
    }
    return upgrades
}
