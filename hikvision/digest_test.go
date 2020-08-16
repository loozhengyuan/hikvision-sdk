package hikvision

import (
	"reflect"
	"testing"
)

func TestHash_basic(t *testing.T) {
	// Create test cases
	cases := []struct {
		name  string
		input string
		hash  string
	}{
		{"RFC2617ReferenceA1", "Mufasa:testrealm@host.com:Circle Of Life", "939e7578ed9e3c518a452acee763bce9"},
		{"RFC2617ReferenceA2", "GET:/dir/index.html", "39aff3a2bab6126f332b942af96d3366"},
	}
	// Run test cases
	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := Hash(tc.input)
			want := tc.hash
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %+v want %+v", got, want)
			}
		})
	}
}

func TestNewChallenge_basic(t *testing.T) {
	// Create test cases
	cases := []struct {
		name      string
		header    string
		challenge *Challenge
	}{
		{
			"RFC2617Reference",
			`Digest
				realm="testrealm@host.com",
				qop="auth",
				nonce="dcd98b7102dd2f0e8b11d0f600bfb0c093",
				opaque="5ccc069c403ebaf9f0171e9517f40e41"`,
			&Challenge{
				Realm:  "testrealm@host.com",
				Qop:    "auth",
				Nonce:  "dcd98b7102dd2f0e8b11d0f600bfb0c093",
				Opaque: "5ccc069c403ebaf9f0171e9517f40e41",
			},
		},
	}
	// Run test cases
	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := NewChallenge(tc.header)
			want := tc.challenge
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %+v want %+v", got, want)
			}
		})
	}
}

func TestChallenge_Authorize_basic(t *testing.T) {
	// Create test cases
	cases := []struct {
		name      string
		challenge *Challenge
		username  string
		password  string
		method    string
		uri       string
		response  string
	}{
		{
			"RFC2617Reference",
			&Challenge{
				Realm:  "testrealm@host.com",
				Qop:    "auth",
				Nonce:  "dcd98b7102dd2f0e8b11d0f600bfb0c093",
				Opaque: "5ccc069c403ebaf9f0171e9517f40e41",
			},
			"Mufasa",
			"Circle Of Life",
			"GET",
			"/dir/index.html",
			"6629fae49393a05397450978507c4ef1",
		},
	}
	// Run test cases
	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.challenge.Authorize(tc.username, tc.password, tc.method, tc.uri)
			want := &Response{
				Username: tc.username,
				Realm:    tc.challenge.Realm,
				Nonce:    tc.challenge.Nonce,
				URI:      tc.uri,
				Qop:      tc.challenge.Qop,
				Cnonce:   "0a4f113b",
				Nc:       "00000001",
				Response: tc.response,
				Opaque:   tc.challenge.Opaque,
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %+v want %+v", got, want)
			}
		})
	}
}

func TestResponse_String_basic(t *testing.T) {
	// Create test cases
	cases := []struct {
		name     string
		response *Response
		header   string
	}{
		{
			"RFC2617Reference",
			&Response{
				Username: "Mufasa",
				Realm:    "testrealm@host.com",
				Nonce:    "dcd98b7102dd2f0e8b11d0f600bfb0c093",
				URI:      "/dir/index.html",
				Qop:      "auth",
				Cnonce:   "0a4f113b",
				Nc:       "00000001",
				Response: "6629fae49393a05397450978507c4ef1",
				Opaque:   "5ccc069c403ebaf9f0171e9517f40e41",
			},
			`Digest username="Mufasa", realm="testrealm@host.com", nonce="dcd98b7102dd2f0e8b11d0f600bfb0c093", uri="/dir/index.html", qop=auth, nc=00000001, cnonce="0a4f113b", response="6629fae49393a05397450978507c4ef1", opaque="5ccc069c403ebaf9f0171e9517f40e41"`,
		},
	}
	// Run test cases
	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.response.String()
			want := tc.header
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
