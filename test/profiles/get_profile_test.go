package test_users

import (
	"context"
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/pkg/client"
	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/test/utils"
)

func TestGivenLoggedInWhenGetProfileShouldReturnProfile(t *testing.T) {
	followeeUsername := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Username())
	followeeEmail := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Email())
	followeePassword := faker.Password()

	followeeClient := client.NewClient()

	_, err := followeeClient.Users.RegisterUser(context.Background(), followeeUsername, followeeEmail, followeePassword)
	if err != nil {
		t.Fatal(err)
	}

	followerUsername := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Username())
	followerEmail := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Email())
	followerPassword := faker.Password()

	followerClient := client.NewClient()

	_, err = followerClient.Users.RegisterUser(context.Background(), followerUsername, followerEmail, followerPassword)
	if err != nil {
		t.Fatal(err)
	}

	err = followerClient.Login(followerEmail, followerPassword)
	if err != nil {
		t.Fatal(err)
	}

	_, err = followerClient.Profiles.FollowUser(context.Background(), followeeUsername)
	if err != nil {
		t.Fatal(err)
	}

	profile, err := followerClient.Profiles.GetProfile(context.Background(), followeeUsername)
	if err != nil {
		t.Fatal(err)
	}

	if profile.Profile.Username != followeeUsername {
		t.Fatalf("got %s, want %s", profile.Profile.Username, followeeUsername)
	}

	if !profile.Profile.Following {
		t.Fatalf("got %t, want %t", profile.Profile.Following, true)
	}
}

func TestGivenNotLoggedInWhenGetProfileShouldReturnProfile(t *testing.T) {
	followeeUsername := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Username())
	followeeEmail := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Email())
	followeePassword := faker.Password()

	followeeClient := client.NewClient()

	_, err := followeeClient.Users.RegisterUser(context.Background(), followeeUsername, followeeEmail, followeePassword)
	if err != nil {
		t.Fatal(err)
	}

	followerUsername := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Username())
	followerEmail := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Email())
	followerPassword := faker.Password()

	followerClient := client.NewClient()

	_, err = followerClient.Users.RegisterUser(context.Background(), followerUsername, followerEmail, followerPassword)
	if err != nil {
		t.Fatal(err)
	}

	profile, err := followerClient.Profiles.GetProfile(context.Background(), followeeUsername)
	if err != nil {
		t.Fatal(err)
	}

	if profile.Profile.Username != followeeUsername {
		t.Fatalf("got %s, want %s", profile.Profile.Username, followeeUsername)
	}

	if profile.Profile.Following {
		t.Fatalf("got %t, want %t", profile.Profile.Following, false)
	}
}

func TestGivenNotFoundWhenGetProfileShouldReturnError(t *testing.T) {
	followeeUsername := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Username())

	followerUsername := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Username())
	followerEmail := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Email())
	followerPassword := faker.Password()

	followerClient := client.NewClient()

	_, err := followerClient.Users.RegisterUser(context.Background(), followerUsername, followerEmail, followerPassword)
	if err != nil {
		t.Fatal(err)
	}

	_, err = followerClient.Profiles.GetProfile(context.Background(), followeeUsername)
	if err == nil {
		t.Fatal("Should have returned an error")
	}
}
