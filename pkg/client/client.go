package client

import (
	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/internal/users"
)

type client struct {
	Users users.UsersClient
}

func NewClient(email string, password string) client {
	baseURL := "https://realworld-example-app-api-api-gateway-4os2xo04.uc.gateway.dev/api"

	usersClient := users.NewUsersClient(baseURL, email, password)

	return client{
		Users: usersClient,
	}
}

func NewClientWithBaseUrl(baseURL string, email string, password string) client {
	usersClient := users.NewUsersClient(baseURL, email, password)

	return client{
		Users: usersClient,
	}
}
