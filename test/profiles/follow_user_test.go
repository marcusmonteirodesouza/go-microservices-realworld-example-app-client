package test_users

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/pkg/client"
	"github.com/marcusmonteirodesouza/go-microservices-realworld-example-app-client/test/utils"
)

func TestGivenValidRequestWhenFollowUserShouldReturnProfile(t *testing.T) {
	followeeUsername := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Username())
	followeeEmail := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Email())
	followeePassword := faker.Password()

	followeeClient := client.NewClient()

	_, err := followeeClient.Users.RegisterUser(followeeUsername, followeeEmail, followeePassword)
	if err != nil {
		t.Fatal(err)
	}

	followerUsername := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Username())
	followerEmail := fmt.Sprintf("%s%s", utils.TestPrefix, faker.Email())
	followerPassword := faker.Password()

	followerClient := client.NewClient()

	_, err = followerClient.Users.RegisterUser(followerUsername, followerEmail, followerPassword)
	if err != nil {
		t.Fatal(err)
	}

	err = followerClient.Login(followerEmail, followerPassword)
	if err != nil {
		t.Fatal(err)
	}

	profile, err := followerClient.Profiles.FollowUser(followeeEmail)

	if profile.Profile.Username != followeeUsername {
		t.Fatalf("got %s, want %s", profile.Profile.Username, followeeUsername)
	}

	if !profile.Profile.Following {
		t.Fatalf("got %t, want %t", profile.Profile.Following, true)
	}
}
