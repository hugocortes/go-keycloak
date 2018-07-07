package keycloak

import (
	"context"
	"fmt"
)

// AuthenticationService ...
type AuthenticationService service

// AccessGrantRequest ...
type AccessGrantRequest struct {
	GrantType    string `url:"grant_type"`
	ClientID     string `url:"client_id"`
	ClientSecret string `url:"client_secret"`
	Scope        string `url:"scope,omitempty"`
	Username     string `url:"username,omitempty"`
	Password     string `url:"password,omitempty"`
}

// OIDCToken ...
type OIDCToken struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not_before_policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

// GetOIDCToken ...
func (c *AuthenticationService) GetOIDCToken(
	ctx context.Context,
	grantReq *AccessGrantRequest,
) (*OIDCToken, *Response, error) {
	path := fmt.Sprintf("%s/%s/protocol/openid-connect/token", defaultBase, c.client.realm)
	h := headers{contentType: formEncoded}

	req, err := c.client.newRequest("POST", path, grantReq, h)
	if err != nil {
		return nil, nil, err
	}

	token := new(OIDCToken)
	resp, err := c.client.do(ctx, req, token)
	if err != nil {
		return nil, resp, err
	}

	return token, resp, nil
}
