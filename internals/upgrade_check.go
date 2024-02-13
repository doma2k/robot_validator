package internals

import (
	"log"
	"os"
	"path/filepath"
)

func GetStagedBinaries() []string {
    // home, err := os.UserHomeDir()
    // if err != nil {
    //     log.Fatal(err)
    // }
    path := filepath.Join("/home/elys/", ".elys/cosmovisor/upgrades")

    upgrades := []string{}
    files, err := os.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }
    for _, file := range files {
        upgrades = append(upgrades, file.Name())
    }
    return upgrades
}
