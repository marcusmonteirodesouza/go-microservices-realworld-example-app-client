package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/internal/common"
)

type UsersClient struct {
	BaseURL string
}

func NewUsersClient(baseURL string) UsersClient {
	return UsersClient{
		BaseURL: baseURL,
	}
}

type user struct {
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

type RegisterUserResponse = user

func newRegisterUserRequest(username string, email string, password string) registerUserRequest {
	return registerUserRequest{
		User: registerUserRequestUser{
			Username: username,
			Email:    email,
			Password: password,
		},
	}
}

func (c *UsersClient) RegisterUser(username string, email string, password string) (*RegisterUserResponse, error) {
	url := fmt.Sprintf("%s/users", c.BaseURL)

	requestData := newRegisterUserRequest(username, email, password)

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusCreated:
		responseData := &RegisterUserResponse{}
		err = json.NewDecoder(response.Body).Decode(&responseData)
		if err != nil {
			return nil, err
		}
		return responseData, nil
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
