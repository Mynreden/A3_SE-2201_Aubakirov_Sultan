package tests

import (
	"io"
	"log"
	"net/http"
)

func printResponseBody(resp *http.Response) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Print(sb)
}

func printRequestBody(req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Print(sb)

}
