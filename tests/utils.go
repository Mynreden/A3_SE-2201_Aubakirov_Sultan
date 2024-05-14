package tests

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func readResponseBody(resp *http.Response) string {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	return sb
}

func parseResponseBody(resp *http.Response) (error, map[string]interface{}) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)

	var data map[string]interface{}
	err = json.Unmarshal([]byte(sb), &data)

	if err != nil {
		return err, nil
	}

	return nil, data
}
