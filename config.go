package keycloak

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Keycloak ...
type Keycloak struct {
	// Configuration
	baseURL      string
	realm        string
	clientID     string
	clientName   string
	clientSecret string
	adminOIDC    *OIDCToken

	// Admin Routes
	adminUsers     string
	adminResources string
	adminEvaluate  string

	// UMA Routes
	umaAuth     string
	umaToken    string
	umaUserInfo string
	umaResource string
}

func (client *Keycloak) setUMATokenPath() {
	client.umaToken = client.baseURL + "/auth/realms/" + client.realm + "/protocol/openid-connect/token"
}

// BaseURL ...
func (client Keycloak) BaseURL() string { return client.baseURL }

// Realm ...
func (client Keycloak) Realm() string { return client.realm }

// ClientID ...
func (client Keycloak) ClientID() string { return client.clientID }

// ClientName ...
func (client Keycloak) ClientName() string { return client.clientName }

// ClientSecret ...
func (client Keycloak) ClientSecret() string { return client.clientSecret }

func doHTTPRequest(req *http.Request) (*json.Decoder, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	status := resp.StatusCode
	if status >= 400 {
		err = fmt.Errorf("%v response", status)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	byteReader := bytes.NewReader(body)
	decoder := json.NewDecoder(byteReader)

	return decoder, nil
}
