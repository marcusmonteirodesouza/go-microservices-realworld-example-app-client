package profiles

import (
	"context"
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

func (c *ProfilesClient) FollowUser(ctx context.Context, username string) (*Profile, error) {
	if c.token == nil {
		return nil, errors.New("Please Login first")
	}

	url := fmt.Sprintf("%s/profiles/%s/follow", c.baseURL, username)

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
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
	case http.StatusCreated:
		responseData := Profile{}
		err = json.NewDecoder(response.Body).Decode(&responseData)
		if err != nil {
			return nil, err
		}
		return &responseData, nil
	default:
		return nil, fmt.Errorf("Unexpected HTTP response code %d", response.StatusCode)
	}
}

func (c *ProfilesClient) UnfollowUser(ctx context.Context, username string) (*Profile, error) {
	if c.token == nil {
		return nil, errors.New("Please Login first")
	}

	url := fmt.Sprintf("%s/profiles/%s/follow", c.baseURL, username)

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
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
		responseData := Profile{}
		err = json.NewDecoder(response.Body).Decode(&responseData)
		if err != nil {
			return nil, err
		}
		return &responseData, nil
	default:
		return nil, fmt.Errorf("Unexpected HTTP response code %d", response.StatusCode)
	}
}

func (c *ProfilesClient) GetProfile(ctx context.Context, username string) (*Profile, error) {
	url := fmt.Sprintf("%s/profiles/%s", c.baseURL, username)

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	if c.token != nil {
		req.Header.Set("authorization", fmt.Sprintf("Bearer %s", *c.token))
	}

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		responseData := Profile{}
		err = json.NewDecoder(response.Body).Decode(&responseData)
		if err != nil {
			return nil, err
		}
		return &responseData, nil
	case http.StatusNotFound:
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
