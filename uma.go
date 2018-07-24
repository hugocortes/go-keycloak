package keycloak

import (
	"context"
	"fmt"
)

// UMAService handles communication with Keycloak UMA
type UMAService service

// GetUMAUser allows user to view their token mappings.
// The provided interface is returned to be decoded on success.
func (c *UMAService) GetUMAUser(
	ctx context.Context,
	token string,
	v interface{},
) (interface{}, *Response, error) {
	path := fmt.Sprintf("%s/%s/protocol/openid-connect/userinfo", defaultBase, c.client.realm)
	h := headers{authorization: token}

	req, err := c.client.newRequest("GET", path, nil, h, false)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
