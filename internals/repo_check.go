package internals

import (
	"io"
	"log"
	"net/http"
	"robot-validator/utils"
	"robot-validator/types"
)

func GetLatestRelease(url string) types.Release {
	var data types.Release
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

	utils.UnmarshalJSON(body, &data)
	return data
}
