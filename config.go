package keycloak

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

// Response ...
type Response struct {
	Response *http.Response
}

// ErrorResponse ...
type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"error_description"`
}

// Error ...
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}

// Client ...
type Client struct {
	common     service      // Reuse struct
	httpClient *http.Client // HTTP client to communicate with keycloak

	// Keycloak Client Configuration
	baseURL      *url.URL
	realm        string
	clientID     string
	clientName   string
	clientSecret string

	// Services
	Authentication *AuthenticationService

	// TODO look into oauth2 library: golang.org/x/oauth2
	adminOIDC *OIDCToken
}

type service struct {
	client *Client
}

type headers struct {
	authorization string
	contentType   string
}

// BaseURL ...
func (c Client) BaseURL() string { return c.baseURL.String() }

// Realm ...
func (c Client) Realm() string { return c.realm }

// ClientID ...
func (c Client) ClientID() string { return c.clientID }

// ClientName ...
func (c Client) ClientName() string { return c.clientName }

// ClientSecret ...
func (c Client) ClientSecret() string { return c.clientSecret }

// newRequest builds a new Keycloak request.
func (c *Client) newRequest(
	method,
	path string,
	body interface{},
	h headers,
) (*http.Request, error) {
	u, err := c.baseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	if h.contentType == formEncoded && body != nil {
		formEnc, err := query.Values(body)
		if err != nil {
			return nil, err
		}
		form := strings.NewReader(formEnc.Encode())
		req, err = http.NewRequest(method, u.String(), form)
	} else if body != nil {
		buf := new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(method, u.String(), buf)
	}
	if err != nil {
		return nil, err
	}

	if h.contentType != "" {
		req.Header.Set("Content-Type", h.contentType)
	}
	if body != nil && h.contentType == "" {
		req.Header.Set("Content-Type", "application/json")
	}
	if h.authorization != "" {
		req.Header.Set("Authorization", h.authorization)
	}

	return req, nil
}

// do ...
func (c *Client) do(
	ctx context.Context,
	req *http.Request,
	v interface{},
) (*Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}
	defer resp.Body.Close()

	response := &Response{Response: resp}

	if c := resp.StatusCode; c >= 300 {
		errorResponse := &ErrorResponse{Response: resp}

		data, err := ioutil.ReadAll(resp.Body)
		if err == nil && data != nil {
			json.Unmarshal(data, errorResponse)
		}

		return nil, errorResponse
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore empty response errors
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return response, err
}
