package keycloak

import (
	"context"
	"fmt"
)

// AdminUserService handles communication with keycloak user management
type AdminUserService service

// User represents the Keycloak user
type User struct {
	Access                     *map[string]interface{} `json:"access,omitempty"`
	Attributes                 *map[string]interface{} `json:"attributes,omitempty"`
	ClientConsents             *[]UserConsent          `json:"clientConsents,omitempty"`
	ClientRoles                *map[string]interface{} `json:"clientRoles,omitempty"`
	CreatedTimestamp           *int64                  `json:"createdTimestamp,omitempty"`
	Credentials                *[]Credential           `json:"credentials,omitempty"`
	DisableableCredentialTypes *[]string               `json:"disableableCredentialTypes,omitempty"`
	Email                      *string                 `json:"email,omitempty"`
	EmailVerified              *bool                   `json:"emailVerified,omitempty"`
	Enabled                    *bool                   `json:"enabled,omitempty"`
	FederatedIdentities        *[]FederatedIdentity    `json:"federatedIdentities,omitempty"`
	FederationLink             *string                 `json:"federationLink,omitempty"`
	FirstName                  *string                 `json:"firstName,omitempty"`
	Groups                     *[]string               `json:"groups,omitempty"`
	ID                         *string                 `json:"id,omitempty"`
	LastName                   *string                 `json:"lastName,omitempty"`
	NotBefore                  *int32                  `json:"notBefore,omitempty"`
	Origin                     *string                 `json:"origin,omitempty"`
	RealmRoles                 *[]string               `json:"realmRoles,omitempty"`
	RequiredActions            *[]string               `json:"requiredActions,omitempty"`
	Self                       *string                 `json:"self,omitempty"`
	ServiceAccountClientID     *string                 `json:"serviceAccountClientId,omitempty"`
	Username                   *string                 `json:"username,omitempty"`
}

// FederatedIdentity represents third party signups
type FederatedIdentity struct {
	IdentityProvider *string `json:"identityProvider,omitempty"`
	UserID           *string `json:"userId,omitempty"`
	UserName         *string `json:"userName,omitempty"`
}

// UserConsent represents scopes that have been consented
type UserConsent struct {
	ClientID               *string                 `json:"clientId,omitempty"`
	CreatedDate            *int64                  `json:"createdDate,omitempty"`
	GrantedClientRoles     *map[string]interface{} `json:"grantedClientRoles,omitempty"`
	GrantedProtocolMappers *map[string]interface{} `json:"grantedProtocolMappers,omitempty"`
	GrantedRealmRoles      *[]string               `json:"grantedRealmRoles,omitempty"`
	LastUpdatedDate        *int64                  `json:"lastUpdatedDate,omitempty"`
}

// Credential represents the user's credentials type
type Credential struct {
	Algorithm         *string             `json:"algorithm,omitempty"`
	Config            *MultivaluedHashMap `json:"config,omitempty"`
	Counter           *int32              `json:"counter,omitempty"`
	CreatedDate       *int64              `json:"createdDate,omitempty"`
	Device            *string             `json:"device,omitempty"`
	Digits            *int32              `json:"digits,omitempty"`
	HashIterations    *int32              `json:"hashIterations,omitempty"`
	HashedSaltedValue *string             `json:"hashedSaltedValue,omitempty"`
	Period            *int32              `json:"period,omitempty"`
	Salt              *string             `json:"salt,omitempty"`
	Temporary         *bool               `json:"temporary,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Value             *string             `json:"value,omitempty"`
}

// MultivaluedHashMap ...
type MultivaluedHashMap struct {
	Empty      *bool  `json:"empty,omitempty"`
	LoadFactor *int32 `json:"loadFactor,omitempty"`
	Threshold  *int32 `json:"threshold,omitempty"`
}

// GetUserByID retrieves a user by ID
func (c *AdminUserService) GetUserByID(
	ctx context.Context,
	ID string,
) (*User, *Response, error) {
	path := fmt.Sprintf("%s/%s/users/%s", defaultAdminBase, c.client.realm, ID)

	req, err := c.client.newRequest("GET", path, nil, headers{}, true)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)
	resp, err := c.client.do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}
