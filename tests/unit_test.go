package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestNewAccountCreationRequest(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Name":     "Ais",
		"Email":    "tabuldin.a@mail.ru",
		"Password": "password123",
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:4000/v1/users", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	printResponseBody(resp)
}

func TestNewAccountCreationRequest2(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Name":     "Ais",
		"Email":    "something not an email",
		"Password": ""})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:4000/v1/users", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	printResponseBody(resp)
}

func TestGettingAccountAuthenticationToken(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Email":    "tabuldin.a@mail.ru",
		"Password": "password123",
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:4000/v1/tokens/authentication", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	printResponseBody(resp)
}

func TestGettingAccountAuthenticationToken2(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Email":    "tabtabtab@tab.a", // This email does not exist
		"Password": "password123",
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:4000/v1/tokens/authentication", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	printResponseBody(resp)
}

func TestActivatingAccount(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Token": "BPNT6BKDNMXT532QU7PTJJCV5Q", // ACTIVATION TOKEN
	})
	responseBody := bytes.NewBuffer(postBody)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, "http://localhost:4000/v1/users/activated", responseBody)
	if err != nil {
		log.Fatalf("An Error Occurred while creating request: %v", err)
	}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("An Error Occurred while sending request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	printResponseBody(resp)
}

func TestActivatingAccount2(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Token": "NOT A VALID TOKEN", // ACTIVATION TOKEN
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.NewRequest(http.MethodPost, "http://localhost:4000/v1/users/activated", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	printRequestBody(resp)
}
