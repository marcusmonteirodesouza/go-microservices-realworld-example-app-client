package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/internal/common"
)

type UsersClient struct {
	baseURL string
	token   *string
}

func NewUsersClient(baseURL string) UsersClient {
	return UsersClient{
		baseURL: baseURL,
	}
}

func NewUsersClientWithToken(baseURL string, token string) UsersClient {
	return UsersClient{
		baseURL: baseURL,
		token:   &token,
	}
}

type User struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
	} `json:"user"`
}

type registerUserRequest struct {
	User registerUserRequestUser `json:"user"`
}

type registerUserRequestUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func newRegisterUserRequest(username string, email string, password string) registerUserRequest {
	return registerUserRequest{
		User: registerUserRequestUser{
			Username: username,
			Email:    email,
			Password: password,
		},
	}
}

type updateUserRequest struct {
	User UpdateUserRequestUser `json:"user"`
}

type UpdateUserRequestUser struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Bio      *string `json:"bio"`
	Image    *string `json:"image"`
}

func (c *UsersClient) RegisterUser(username string, email string, password string) (*User, error) {
	url := fmt.Sprintf("%s/users", c.baseURL)

	requestData := newRegisterUserRequest(username, email, password)

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusCreated:
		responseData := User{}
		err = json.NewDecoder(response.Body).Decode(&responseData)
		if err != nil {
			return nil, err
		}
		return &responseData, nil
	case http.StatusUnprocessableEntity:
		errorResponse := &common.ErrorResponse{}
		err = json.NewDecoder(response.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", errorResponse.Errors.Body[0])
	default:
		return nil, fmt.Errorf("Unexpected HTTP response code %d", response.StatusCode)
	}
}

func (c *UsersClient) GetCurrentUser() (*User, error) {
	if c.token == nil {
		return nil, errors.New("Please Login first")
	}

	url := fmt.Sprintf("%s/user", c.baseURL)
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", *c.token))

	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		responseData := User{}
		err = json.NewDecoder(response.Body).Decode(&responseData)
		if err != nil {
			return nil, err
		}
		return &responseData, nil
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("Unauthorized")
	default:
		return nil, fmt.Errorf("Unexpected HTTP response code %d", response.StatusCode)
	}
}

func (c *UsersClient) UpdateUser(request UpdateUserRequestUser) (*User, error) {
	if c.token == nil {
		return nil, errors.New("Please Login first")
	}

	url := fmt.Sprintf("%s/user", c.baseURL)
	client := &http.Client{}

	requestBody := updateUserRequest{
		User: request,
	}
	requestBytes, err := json.Marshal(requestBody)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", *c.token))

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		responseData := &User{}
		err = json.NewDecoder(response.Body).Decode(&responseData)
		if err != nil {
			return nil, err
		}
		return responseData, nil
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("Unauthorized")
	case http.StatusUnprocessableEntity:
		errorResponse := &common.ErrorResponse{}
		err = json.NewDecoder(response.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", errorResponse.Errors.Body[0])
	default:
		return nil, fmt.Errorf("Unexpected HTTP response code %d", response.StatusCode)
	}
}

func (c *UsersClient) SetToken(token string) {
	c.token = &token
}
