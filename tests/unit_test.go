package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"testing"

	_ "github.com/stretchr/testify/assert"
)

type User struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Activated bool   `json:"activated"`
}

var testToken string

func TestAccountCreationRequest(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Name":     "Sultan",
		"Email":    "sultan.a@gmail.ru",
		"Password": "password",
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

	var response struct {
		User User `json:"user"`
	}

	str := readResponseBody(resp)

	if err := json.Unmarshal([]byte(str), &response); err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := json.Unmarshal([]byte(str), &response); err != nil {
		fmt.Println("Error:", err)
		assert.Fail(t, "Error due Response Unmarshall")
		return
	}
	assert.Positive(t, response.User.ID)
	assert.Equal(t, "Sultan", response.User.Name)
	assert.Equal(t, "sultan.a@gmail.ru", response.User.Email)
}

type ErrorIncorrectAccountCreationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TestIncorrectAccountCreationRequest(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Name":     "Sultan",
		"Email":    "Wrong Email",
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

	var response struct {
		Error ErrorIncorrectAccountCreationRequest `json:"error"`
	}

	str := readResponseBody(resp)

	if err := json.Unmarshal([]byte(str), &response); err != nil {
		fmt.Println("Error:", err)
		return
	}
	assert.Equal(t, "must be provided", response.Error.Password)
	assert.Equal(t, "must be valid email address", response.Error.Email)
}

type TokenForTestLogin struct {
	Token  string `json:"token"`
	Expiry string `json:"expiry"`
}

func TestLogin(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Email":    "sultan.a@gmail.ru",
		"Password": "password",
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

	var response struct {
		Token TokenForTestLogin `json:"authentication_token"`
	}

	str := readResponseBody(resp)

	if err := json.Unmarshal([]byte(str), &response); err != nil {
		fmt.Println("Error:", err)
		return
	}
	testToken = response.Token.Token
	assert.NotNil(t, response.Token.Token)
	assert.NotNil(t, response.Token.Expiry)
}

func TestIncorrectLogin(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Email":    "tabtabtab@tab.a", // This email does not exist
		"Password": "password",
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

	var response struct {
		Error string `json:"error"`
	}

	str := readResponseBody(resp)

	if err := json.Unmarshal([]byte(str), &response); err != nil {
		fmt.Println("Error:", err)
		return
	}
	assert.Equal(t, "invalid authentication credentials", response.Error)
}

func TestValidToken(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Token": testToken, // ACTIVATION TOKEN
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

	str := readResponseBody(resp)

	var response struct {
		Error struct {
			Token string `json:"token"`
		} `json:"error"`
	}

	if err := json.Unmarshal([]byte(str), &response); err != nil {
		fmt.Println("Error:", err)
		return
	}
	assert.Equal(t, "invalid or expired activation token", response.Error.Token)
}

func TestInvalidToken(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Token": "NOT A VALID TOKEN", // ACTIVATION TOKEN
	})
	responseBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:4000/v1/users/activated", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var response struct {
		Error string `json:"error"`
	}

	str := readResponseBody(resp)
	if err := json.Unmarshal([]byte(str), &response); err != nil {
		fmt.Println("Error:", err)
		return
	}
	assert.Equal(t, "rate limited exceeded", response.Error)
}
