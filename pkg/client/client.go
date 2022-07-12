package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/internal/profiles"
	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/internal/users"
)

type Client struct {
	baseURL  string
	Users    users.UsersClient
	Profiles profiles.ProfilesClient
}

var baseURL = "https://realworld-example-app-api-gateway-36qfdayk.uc.gateway.dev"

func NewClient() Client {
	return Client{
		baseURL:  baseURL,
		Users:    users.NewUsersClient(baseURL),
		Profiles: profiles.NewProfilesClient(baseURL),
	}
}

func NewClientWithBaseUrl(baseURL string) Client {
	return Client{
		baseURL:  baseURL,
		Users:    users.NewUsersClient(baseURL),
		Profiles: profiles.NewProfilesClient(baseURL),
	}
}

func NewClientWithToken(token string) Client {
	return Client{
		baseURL:  baseURL,
		Users:    users.NewUsersClientWithToken(baseURL, token),
		Profiles: profiles.NewProfilesClientWithToken(baseURL, token),
	}
}

func NewClientWithBaseUrlAndToken(baseURL string, token string) Client {
	return Client{
		baseURL:  baseURL,
		Users:    users.NewUsersClientWithToken(baseURL, token),
		Profiles: profiles.NewProfilesClientWithToken(baseURL, token),
	}
}

type loginRequest struct {
	User loginRequestUser `json:"user"`
}

type loginRequestUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func newLoginRequest(email string, password string) loginRequest {
	return loginRequest{
		User: loginRequestUser{
			Email:    email,
			Password: password,
		},
	}
}

func (c *Client) Login(email string, password string) error {
	url := fmt.Sprintf("%s/users/login", c.baseURL)

	requestData := newLoginRequest(email, password)

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		return err
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		responseData := users.User{}
		err = json.NewDecoder(response.Body).Decode(&responseData)
		if err != nil {
			return err
		}
		c.Users.SetToken(responseData.User.Token)
		c.Profiles.SetToken(responseData.User.Token)
		return nil
	case http.StatusUnauthorized:
		return fmt.Errorf("Unauthorized")
	default:
		return fmt.Errorf("Unexpected HTTP response code %d", response.StatusCode)
	}
}
