# go-keycloak

go-keycloak is a Go client library for accessing the [Keycloak API](https://www.keycloak.org/documentation.html)

## Usage

```go
import "github.com/hugocortes/go-keycloak"
```

Constructing the Keycloak client depends on the client that will be used to make requests and if that user or client has offline access to disable the SSO idle timeout. This provides flexibiliy in creating more than one Keycloak client to authenticate against different realms and/or clients.

1. Using a Service Account will require the client ID, client name, and the client secret
```go
// Creates a service account
serviceAccount := keycloak.NewServiceAccount(
		httpClient, // httpClient or use default if nil
		"BASE_URL", // base keycloak url
    "REALM", // target realm
    hasOfflineAccess, // If offline_access role is assigned
		"CLIENT_ID", // target client id
		"CLIENT_NAME", // target client name
		"CLIENT_SECRET", // target client secret
	)
```
2. Using a user to authenticate using a confidential client will require client ID, client name, client secret, admin, and admin password
```go
// Creates a service account
serviceAccount := keycloak.NewConfidentialAdmin(
		httpClient, // httpClient or use default if nil
		"BASE_URL", // base keycloak url
    "REALM", // target realm
    hasOfflineAccess, // If offline_access role is assigned
		"CLIENT_ID", // target client id
		"CLIENT_NAME", // target client name
    "CLIENT_SECRET", // target client secret
    "ADMION_USER", // target admin username
    "ADMIN_PASS", // target admin password
	)
```
3. User a user to authenticate using a public client will require client ID, client name, admin, and admin password
```go
// Creates a service account
serviceAccount := keycloak.NewPublicAdmin(
		httpClient, // httpClient or use default if nil
		"BASE_URL", // base keycloak url
    "REALM", // target realm
    hasOfflineAccess, // If offline_access role is assigned
		"CLIENT_ID", // target client id
		"CLIENT_NAME", // target client name
    "ADMION_USER", // target admin username
    "ADMIN_PASS", // target admin password
	)
```

Note: Depending on the type of request, the library will require the Client (if full scope mapping is disbled) and Admin User and/or Service Account to have the appropriate role(s) or 403 errors will be returned.
