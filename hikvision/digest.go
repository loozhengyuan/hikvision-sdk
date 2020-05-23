package hikvision

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

// Challenge represents the digest challenge,
// which is found on the `WWW-Authenticate`
// header on the initial request.
type Challenge struct {
	Realm     string
	Domain    string
	Nonce     string
	Opaque    string
	Stale     string
	Algorithm string
	Qop       string // TODO: Should Qop be an array of strings?
}

// Response represents the digest response,
// which is embedded on the `Authorization`
// header on every subsequent requests.
type Response struct {
	Username string
	Realm    string
	Nonce    string
	URI      string
	Qop      string
	Cnonce   string
	Nc       string
	Response string
	Opaque   string
}

// Hash returns a MD5 hash of a string.
func Hash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// NewChallenge is a constructor for Challenge
func NewChallenge(header string) *Challenge {
	// Parse digest
	re := regexp.MustCompile(`(\w+)="([^"]*)"`)
	arr := re.FindAllStringSubmatch(header, -1)

	// Create digest
	m := map[string]string{}
	for _, a := range arr {
		k := a[1]
		v := a[2]
		m[k] = v
	}

	// Marshall digest
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	// Unmarshall into struct
	challenge := Challenge{}
	if err := json.Unmarshal(b, &challenge); err != nil {
		log.Fatalln("Error:", err)
	}
	return &challenge
}

// Authorize method returns a Response.
func (c *Challenge) Authorize(username, password, method, uri string) *Response {
	// Get A1
	a1 := Hash(fmt.Sprintf("%s:%s:%s", username, c.Realm, password))

	// Get A2
	a2 := Hash(fmt.Sprintf("%s:%s", method, uri))

	// Get full response
	response := Hash(fmt.Sprintf("%s:%s:%s:%s:%s:%s", a1, c.Nonce, "00000001", "0a4f113b", c.Qop, a2))

	// Create response
	r := &Response{
		Username: username,
		Realm:    c.Realm,
		Nonce:    c.Nonce,
		URI:      uri,
		Qop:      c.Qop,
		Cnonce:   "0a4f113b",
		Nc:       "00000001",
		Response: response,
		Opaque:   c.Opaque,
	}
	return r
}

// Header method returns Response as header string
func (r *Response) String() string {
	// TODO: Why isn't qop and nc quoted in RFC2617? Is it supposed to be case?
	return fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", qop=%s, nc=%s, cnonce="%s", response="%s", opaque="%s"`, r.Username, r.Realm, r.Nonce, r.URI, r.Qop, r.Nc, r.Cnonce, r.Response, r.Opaque)
}
