package other

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Dog struct {
	Message string `json:"message"`
	Status  string `json:"status`
}

// This function's purpose is to show a simple REST consumer with api-key auth
// TODO: Pull API key out into environment variable
func GetDogBreedPic(download bool) Dog {
	// Function alias for convenience
	_logf := log.Printf

	url := "https://dog.ceo/api/breeds/image/random"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		_logf("Could not prepare our new request - %v\n", err)
	}

	// !!! Do not hard-code api-keys. This is a test api, so we don't care.
	// !!! IN ANY PRODUCTION CASE, DO NOT HARD-CODE SECRETS!!
	req.Header.Add("Authorization", "api-key afgl6752390nasdjghtuyvbn786")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		_logf("Could not perform our new request - %v\n", err)
	}
	bodyBytes, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		_logf("Could not read our response body - %v\n", err)
	}

	var dog Dog
	if err := json.Unmarshal(bodyBytes, &dog); err != nil {
		_logf("Could not unmarshall our response body - %v\n", err)
	}

	if download {
		dogPic, err := http.Get(string(dog.Message))
		if err != nil {
			_logf("Could not download dog pic - %v\n", err)
		}
		defer dogPic.Body.Close()

		dogFile, err := os.Create("dog.jpg")
		if err != nil {
			_logf("Could not create output file - %v\n", err)
		}
		defer dogFile.Close()

		_, err = io.Copy(dogFile, dogPic.Body)
		if err != nil {
			_logf("Could not copy body to output file - %v\n", err)
		}
	}
	return dog

}
