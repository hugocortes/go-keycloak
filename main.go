package keycloak

// New keycloak clients
func New(
	baseURL string,
	realm string,
	clientID string,
	clientName string,
	clientSecret string,
) *Keycloak {

	keycloakClient := &Keycloak{
		baseURL:      baseURL,
		realm:        realm,
		clientID:     clientID,
		clientName:   clientName,
		clientSecret: clientSecret,
		adminOIDC:    &OIDCToken{},
	}

	keycloakClient.setUMATokenPath()

	return keycloakClient
}
