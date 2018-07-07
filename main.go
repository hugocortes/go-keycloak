package keycloak

import (
	"net/http"
	"net/url"
)

const (
	defaultAdminBase = "auth/admin/realms"
	defaultBase      = "auth/realms"

	formEncoded = "application/x-www-form-urlencoded"
)

// NewClient keycloak clients
func NewClient(
	httpClient *http.Client,

	baseURL string,
	realm string,
	clientID string,
	clientName string,
	clientSecret string,
) *Client {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	base, _ := url.Parse(baseURL)

	c := &Client{
		httpClient:   httpClient,
		baseURL:      base,
		realm:        realm,
		clientID:     clientID,
		clientName:   clientName,
		clientSecret: clientSecret,

		// TODO look into oauth2 library: golang.org/x/oauth2
		adminOIDC: &OIDCToken{},
	}

	c.common.client = c
	c.Authentication = (*AuthenticationService)(&c.common)

	return c
}
