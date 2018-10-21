// Package main provides an example for using an admin account or
// a service account to authorize against a client and query a user
// provided a 'query-users' role
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

	serviceAccount := keycloak.NewServiceAccount(
		httpClient,
		os.Getenv("BASE_URL"),
		os.Getenv("REALM"),
		true,
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
	)

	confidentialAdmin := keycloak.NewConfidentialAdmin(
		httpClient,
		os.Getenv("BASE_URL"),
		os.Getenv("REALM"),
		true,
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
		os.Getenv("ADMIN_USER"),
		os.Getenv("ADMIN_PASS"),
	)

	publicAdmin := keycloak.NewPublicAdmin(
		httpClient,
		os.Getenv("BASE_URL"),
		os.Getenv("REALM"),
		true,
		os.Getenv("PUBLIC_CLIENT_ID"),
		os.Getenv("ADMIN_USER"),
		os.Getenv("ADMIN_PASS"),
	)

	fmt.Println("Validating service acount:")
	validate(serviceAccount)
	fmt.Println("Validating confidential admin:")
	validate(confidentialAdmin)
	fmt.Println("Validating public admin:")
	validate(publicAdmin)
}

func validate(kc *keycloak.Client) {
	token, resp, err := kc.Authentication.GetOIDCToken(
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
	fmt.Printf("user token request")
	fmt.Printf("status code: %d \n", resp.Response.StatusCode)
	fmt.Printf("token: %s \n", token.AccessToken)

	// Get the user's token mapping by issuing a request with user token
	reader, resp, err := kc.UMA.GetUMAUser(
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
	user, resp, err := kc.AdminUser.GetUserByID(
		context.Background(),
		userID,
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("User name: %v\n", *user.Username)
}
