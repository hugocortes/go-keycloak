// Package main provides an example for authorizating a user, reading their
// token mappings, and using the client to query the user
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	keycloak "github.com/hugocortes/go-keycloak"
	"github.com/joho/godotenv"
)

var ctx = context.Background()

// UserInfo represents preconfigured mappers for given client
type UserInfo struct {
	Sub string `json:"sub"`
}

// init loads the .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load ENV file %s", err)
	}
}

func main() {
	httpClient := &http.Client{}

	// Create the new keycloak client
	keycloakClient := keycloak.NewClient(
		httpClient,
		os.Getenv("BASE_URL"),
		os.Getenv("REALM"),
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_NAME"),
		os.Getenv("CLIENT_SECRET"),
	)

	// Get the user's password token
	token, resp, err := keycloakClient.Authentication.GetOIDCToken(
		context.Background(),
		&keycloak.AccessGrantRequest{
			GrantType: "password",
			Username:  os.Getenv("EXAMPLE_USERNAME"),
			Password:  os.Getenv("EXAMPLE_PASSWORD"),
		},
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("status code: %d \n", resp.Response.StatusCode)
	fmt.Printf("token: %s \n", token.AccessToken)

	// Get the user's token mapping by issuing a request with user token
	reader, resp, err := keycloakClient.UMA.GetUMAUser(
		context.Background(),
		"Bearer "+token.AccessToken,
		new(UserInfo),
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// GetUMAUser returns an interface that should be decoded to match our struct
	userInfo, ok := reader.(*UserInfo)
	if !ok {
		fmt.Println("error")
		os.Exit(1)
	}

	userID := userInfo.Sub
	fmt.Printf("User ID: %s\n", userID)

	// Issue request using admin token
	user, resp, err := keycloakClient.AdminUser.GetUserByID(
		context.Background(),
		userID,
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("User name: %v\n", *user.Username)
}
