package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"
)

type Movie struct {
	Title   string   `json:"title"`
	Year    int      `json:"year"`
	Runtime string   `json:"runtime"`
	Genres  []string `json:"genres"`
}

func TestInsertingMoviesIntoDatabase(t *testing.T) {
	moviePayload := Movie{
		Title:   "Inception",
		Year:    2010,
		Runtime: "144 mins",
		Genres:  []string{"thriller"},
	}

	payloadBytes, err := json.Marshal(moviePayload)
	if err != nil {
		log.Fatalf("Error marshaling movie payload: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	bearerToken := "ZK633D2HGQGKZSCHQGWRZOLSYI"
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	printResponseBody(resp)
}

func TestInsertingMoviesIntoDatabaseWithWrongYear(t *testing.T) {
	moviePayload := Movie{
		Title:   "Inception",
		Year:    2029, // This is a future date
		Runtime: "144 mins",
		Genres:  []string{"thriller"},
	}

	payloadBytes, err := json.Marshal(moviePayload)
	if err != nil {
		log.Fatalf("Error marshaling movie payload: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	bearerToken := "ZK633D2HGQGKZSCHQGWRZOLSYI"
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	printResponseBody(resp)
}

func TestInsertingMoviesIntoDatabaseWithWrongRuntime(t *testing.T) {
	moviePayload := Movie{
		Title:   "Inception",
		Year:    2020,
		Runtime: "144", // This is not a valid runtime
		Genres:  []string{"thriller"},
	}

	payloadBytes, err := json.Marshal(moviePayload)
	if err != nil {
		log.Fatalf("Error marshaling movie payload: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	bearerToken := "ZK633D2HGQGKZSCHQGWRZOLSYI"
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	printResponseBody(resp)
}

func TestMovieDeletionById(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:4000/v1/movies/5", nil)
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	bearerToken := "ZK633D2HGQGKZSCHQGWRZOLSYI"
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	printResponseBody(resp)
}
