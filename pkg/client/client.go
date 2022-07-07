package client

import "github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/internal/users"

type Client struct {
	Users users.UsersClient
}

func NewClient() Client {
	baseURL := "https://realworld-example-app-api-api-gateway-4os2xo04.uc.gateway.dev/api"

	usersClient := users.NewUsersClient(baseURL)

	return Client{
		Users: usersClient,
	}
}

func NewClientWithBaseUrl(baseURL string) Client {
	usersClient := users.NewUsersClient(baseURL)

	return Client{
		Users: usersClient,
	}
}

func NewClientWithBaseUrlAndToken(baseURL string, token string) Client {
	usersClient := users.NewUsersClientWithToken(baseURL, token)

	return Client{
		Users: usersClient,
	}
}
