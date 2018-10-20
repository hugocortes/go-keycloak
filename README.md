# go-keycloak

go-keycloak is a Go client library for accessing the [Keycloak API](https://www.keycloak.org/documentation.html)

## Usage

```go
import "github.com/hugocortes/go-keycloak"
```

Constructing the Keycloak client depends on the client that will be used to make requests:
1. Using a Service Account will require the client ID, client name, and the client secret.
2. Using a user to authenticate using a confidential client will require client ID, client name, client secret, admin, and admin password
3. User a user to authenticate using a public client will require client ID, client name, admin, and admin password;

This provides flexibiliy in creating more than one Keycloak client to authenticate against different realms and/or clients.

Example:
```go
// Creates a service account
serviceAccount := keycloak.NewServiceAccount(
		httpClient,
		os.Getenv("BASE_URL"),
		os.Getenv("REALM"),
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_NAME"),
		os.Getenv("CLIENT_SECRET"),
	)
```

Note: Verify the account being used has sufficient role access based on the resource query being made. For example, querying users will need the `query-users` role for the account being used.
