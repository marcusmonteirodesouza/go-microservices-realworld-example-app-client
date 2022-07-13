package test_users

import (
	"context"
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/pkg/client"
	api_client "github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/pkg/client"
	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/test/utils"
)

func TestGivenValidRequestWhenUpdateUserShouldReturnUser(t *testing.T) {
	username := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Username())
	email := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Email())
	password := faker.Password()

	client := client.NewClient()

	_, err := client.Users.RegisterUser(context.Background(), username, email, password)
	if err != nil {
		t.Fatal(err)
	}

	usernameUpdate := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Username())
	emailUpdate := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Email())
	bioUpdate := faker.Paragraph()
	imageUpdate := faker.URL()

	updateUserRequest := api_client.UpdateUserRequest{
		Username: &usernameUpdate,
		Email:    &emailUpdate,
		Bio:      &bioUpdate,
		Image:    &imageUpdate,
	}

	err = client.Login(email, password)
	if err != nil {
		t.Fatal(err)
	}

	updatedUser, err := client.Users.UpdateUser(context.Background(), updateUserRequest)
	if err != nil {
		t.Fatal(err)
	}

	if updatedUser.User.Username != *updateUserRequest.Username {
		t.Fatalf("got %s, want %s", updatedUser.User.Username, *updateUserRequest.Username)
	}

	if updatedUser.User.Email != *updateUserRequest.Email {
		t.Fatalf("got %s, want %s", updatedUser.User.Email, *updateUserRequest.Email)
	}

	if updatedUser.User.Bio != *updateUserRequest.Bio {
		t.Fatalf("got %s, want %s", updatedUser.User.Bio, *updateUserRequest.Bio)
	}

	if updatedUser.User.Image != *updateUserRequest.Image {
		t.Fatalf("got %s, want %s", updatedUser.User.Image, *updateUserRequest.Image)
	}
}
