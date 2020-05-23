package hikvision

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// AuthTransport is a http.RoundTripper that takes care of
// authenticating with the API using HTTP Digest Authentication.
type AuthTransport struct {
	Username  string
	Password  string
	Transport http.RoundTripper
}

// NewAuthTransport is a constructor for AuthTransport.
func NewAuthTransport(username, password string) *AuthTransport {
	t := &AuthTransport{
		Username: username,
		Password: password,
	}
	t.Transport = http.DefaultTransport
	return t
}

// RoundTrip implements the RoundTripper interface. It makes a request expecting
// a HTTP 401 Unauthorized response. If unauthenticated, it creates the required
// credentials and follow with the original request.
func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clones the request
	req2 := new(http.Request)
	*req2 = *req

	// Clone request headers
	req2.Header = make(http.Header)
	for k, s := range req.Header {
		req2.Header[k] = s
	}

	// Clone request body
	if req.Body != nil {
		buf, _ := ioutil.ReadAll(req.Body)
		req.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		req2.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	}

	// Make initial request
	resp, err := t.Transport.RoundTrip(req)
	if err != nil || resp.StatusCode != 401 {
		return resp, err
	}

	// Parse digest challenge
	authChal := resp.Header.Get("WWW-Authenticate")
	c := NewChallenge(authChal)

	// Generate digest response
	authResp := c.Authorize(t.Username, t.Password, req.Method, req.RequestURI)

	// We'll no longer use the initial response, so close it
	resp.Body.Close()

	// Make authenticated request.
	req2.Header.Set("Authorization", authResp.String())
	return t.Transport.RoundTrip(req2)
}
