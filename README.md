# go-keycloak

go-keycloak is a Go client library for accessing the [Keycloak API](https://www.keycloak.org/documentation.html)

# Usage

```go
import "github.com/hugocortes/go-keycloak"
```

You will have to construct a new Keycloak client in order to use the different parts of Keycloak.

The following will be required:
* Keycloak base URL
* Realm where your keycloak client was created in
* Client ID
* Client Name
* Client Secret

This provides flexibiliy in creating more than one Keycloak client to authenticate against different realms and/or clients.

Example:
```go
keycloakClient := keycloak.NewClient(
  nil,
  "http://localhost:8080,
  "master",
  "{UUID}",
  "my-client",
  "{UUID}"
)
```

NOTE: When making requests that require elevated roles, this client makes requests with a `client_credential` token which will require Service Accounts Enabled to be true along with appropriate realm-management roles associated with the Service Account. i.e., to query users you will need `query-users` role. 
