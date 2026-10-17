package client

import (
	"encoding/json"
	"io"
	errorstruct "myRoyale/Errors/ErrorStruct"
	"net/http"
	"os"
)

func Fetch(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	apiToken := os.Getenv("CR_API_TOKEN")
	req.Header.Set("Authorization", "Bearer "+apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode == http.StatusOK {
		return bodyBytes, nil
	}

	var statErr errorstruct.StatCode

	json.Unmarshal(bodyBytes, &statErr)
	statErr.Code = resp.StatusCode
	statErr.Message = http.StatusText(resp.StatusCode)

	return []byte{}, &statErr
}
