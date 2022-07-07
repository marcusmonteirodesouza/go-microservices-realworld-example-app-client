package test_users

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/pkg/client"
	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/test/utils"
)

func TestGivenValidRequestWhenGetCurrentUserShouldReturnGetCurrentUserResponse(t *testing.T) {
	username := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Username())
	email := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Email())
	password := faker.Password()

	client := client.NewClient()

	user, err := client.Users.RegisterUser(username, email, password)
	if err != nil {
		t.Fatal(err)
	}

	currentUser, err := client.Users.GetCurrentUser()

	if currentUser.User.Username != username {
		t.Fatalf("got %s, want %s", user.User.Username, username)
	}

	if currentUser.User.Email != email {
		t.Fatalf("got %s, want %s", user.User.Email, email)
	}
}
