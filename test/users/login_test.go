package test_users

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/pkg/client"
	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/test/utils"
)

func TestGivenValidRequestWhenLoginShouldReturnLoginResponse(t *testing.T) {
	username := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Username())
	email := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Email())
	password := faker.Password()

	client := client.NewClient()

	user, err := client.Users.RegisterUser(username, email, password)
	if err != nil {
		t.Fatal(err)
	}

	loggedUser, err := client.Users.Login(email, password)

	if loggedUser.User.Username != username {
		t.Fatalf("got %s, want %s", user.User.Username, username)
	}

	if loggedUser.User.Email != email {
		t.Fatalf("got %s, want %s", user.User.Email, email)
	}
}

func TestGivenUnauthorizedStatusCodeWhenLoginShouldReturnError(t *testing.T) {
	username := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Username())
	email := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Email())
	password := faker.Password()

	client := client.NewClient()

	_, err := client.Users.RegisterUser(username, email, password)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Users.Login(email, faker.Password())

	if err == nil {
		t.Fatal("Should have returned an error")
	}

	if err.Error() != "Unauthorized" {
		t.Fatalf("got %s, want %s", err.Error(), "Unauthorized")
	}
}
