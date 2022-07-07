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

	client := client.NewClient(email, password)

	user, err := client.Users.RegisterUser(username)
	if err != nil {
		t.Fatal(err)
	}

	loggedUser, err := client.Users.Login()

	if loggedUser.User.Username != username {
		t.Fatalf("got %s, want %s", user.User.Username, username)
	}

	if loggedUser.User.Email != email {
		t.Fatalf("got %s, want %s", user.User.Email, email)
	}
}
