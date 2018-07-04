package keycloak

import (
	"net/http"
	"net/url"
	"strings"
)

// HandlePasswordGrant ...
func (client *Keycloak) HandlePasswordGrant(
	username string,
	password string,
	scope string,
) (*OIDCToken, error) {

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Add("username", username)
	data.Add("password", password)
	data.Add("client_id", client.clientName)
	data.Add("client_secret", client.clientSecret)

	if scope != "" {
		data.Add("scope", scope)
	}

	req, _ := http.NewRequest("POST", client.umaToken, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	decoder, err := doHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	response := &OIDCToken{}
	err = decoder.Decode(response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetClientCredentials ...
func (client *Keycloak) GetClientCredentials() (*OIDCToken, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Add("client_id", client.clientName)
	data.Add("client_secret", client.clientSecret)
	data.Add("scope", "offline_access")

	req, _ := http.NewRequest("POST", client.umaToken, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	decoder, err := doHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	response := &OIDCToken{}
	err = decoder.Decode(response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
