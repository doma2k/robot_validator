package main

import (
	"log"
	"net/http"
	"io"
	"encoding/json"
)

func getLatestRelease(url string) Release {
    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
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

    var data Release
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Fatal(err)
    }
    return data
}