package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/internal/common"
)

type UsersClient struct {
	baseURL  string
	email    string
	password string
	token    *string
}

func NewUsersClient(baseURL string, email string, password string) UsersClient {
	return UsersClient{
		baseURL:  baseURL,
		email:    email,
		password: password,
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

type loginRequest struct {
	User loginRequestUser `json:"user"`
}

type loginRequestUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse = user

type GetCurrentUserResponse = user

func newLoginRequest(email string, password string) loginRequest {
	return loginRequest{
		User: loginRequestUser{
			Email:    email,
			Password: password,
		},
	}
}

func (c *UsersClient) RegisterUser(username string) (*RegisterUserResponse, error) {
	url := fmt.Sprintf("%s/users", c.baseURL)

	requestData := newRegisterUserRequest(username, c.email, c.password)

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

func (c *UsersClient) Login() (*RegisterUserResponse, error) {
	url := fmt.Sprintf("%s/users/login", c.baseURL)

	requestData := newLoginRequest(c.email, c.password)

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
	case http.StatusOK:
		responseData := &RegisterUserResponse{}
		err = json.NewDecoder(response.Body).Decode(&responseData)
		if err != nil {
			return nil, err
		}
		return responseData, nil
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("Unauthorized")
	default:
		return nil, fmt.Errorf("Unexpected HTTP response code %d", response.StatusCode)
	}
}

func (c *UsersClient) GetCurrentUser() (*GetCurrentUserResponse, error) {
	url := fmt.Sprintf("%s/user", c.baseURL)
	client := &http.Client{}

	token, err := c.GetToken()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", *token))

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		responseData := &GetCurrentUserResponse{}
		err = json.NewDecoder(response.Body).Decode(&responseData)
		if err != nil {
			return nil, err
		}
		return responseData, nil
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("Unauthorized")
	default:
		return nil, fmt.Errorf("Unexpected HTTP response code %d", response.StatusCode)
	}
}

func (c *UsersClient) GetToken() (*string, error) {
	if c.token != nil {
		return c.token, nil
	}

	loggedUser, err := c.Login()
	if err != nil {
		return nil, err
	}

	c.token = &loggedUser.User.Token
	return c.token, nil
}
