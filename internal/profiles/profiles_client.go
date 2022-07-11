package profiles

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/internal/common"
)

type ProfilesClient struct {
	baseURL string
	token   *string
}

func NewProfilesClient(baseURL string) ProfilesClient {
	return ProfilesClient{
		baseURL: baseURL,
	}
}

func NewProfilesClientWithToken(baseURL string, token string) ProfilesClient {
	return ProfilesClient{
		baseURL: baseURL,
		token:   &token,
	}
}

type Profile struct {
	Profile struct {
		Username  string `json:"username"`
		Bio       string `json:"bio"`
		Image     string `json:"image"`
		Following bool   `json:"following"`
	} `json:"profile"`
}

func (c *ProfilesClient) FollowUser(username string) (*Profile, error) {
	if c.token == nil {
		return nil, errors.New("Please Login first")
	}

	url := fmt.Sprintf("%s/profiles/%s/follow", c.baseURL, username)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", *c.token))

	response, err := http.Post(url, "application/json", nil)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusCreated:
		responseData := Profile{}
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

func (c *ProfilesClient) SetToken(token string) {
	c.token = &token
}
