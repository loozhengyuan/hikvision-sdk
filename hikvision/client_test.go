package hikvision

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestClient_Get(t *testing.T) {
	// Create test cases
	cases := []struct {
		name   string
		status int
		body   string
		err    error
	}{
		{"http200", http.StatusOK, "OK", nil},

		// TODO: Handle 401 Unauthorized
		// TODO: Handle 404 Not Found
		// TODO: Handle 403 Forbidden

		// Returns error for non-2xx http status codes
		{"http301_responseNotOk", http.StatusPermanentRedirect, `{"message":"permanent redirect"}`, ErrResponseNotOk},
		{"http400_responseNotOk", http.StatusBadRequest, `{"message":"bad request"}`, ErrResponseNotOk},

		// Returns error when parsing error messages for non-2xx http status codes
		{"http301_parseErrorMessageFailure", http.StatusPermanentRedirect, `"message":"permanent redirect"`, ErrParseErrorMessageFailure},
		{"http400_parseErrorMessageFailure", http.StatusBadRequest, `{message: bad request}`, ErrParseErrorMessageFailure},
	}

	// Run test cases
	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Mock HTTP server
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", `application/json; charset="UTF-8"`)
				w.WriteHeader(tc.status)
				w.Write([]byte(tc.body))
			})
			server := httptest.NewServer(handler)

			// Parse URL
			u, err := url.Parse(server.URL)
			if err != nil {
				t.Fatalf("error parsing url: %v", err)
			}

			// Create client object
			c, err := NewClient("", "", "")
			if err != nil {
				t.Fatalf("error creating client: %v", err)
			}

			// Execute request
			got, err := c.Get(u)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected error '%v' but got: %v", tc.err, err)
			}

			// Assert response body if non-error responses
			want := tc.body
			if got != nil && string(got) != want {
				t.Errorf("got %+v want %+v", string(got), want)
			}
		})
	}
}

func TestClient_Get_HandleErrorResponses(t *testing.T) {
	// Create test cases
	cases := []struct {
		name        string
		status      int
		contentType string
		fixture     string
		err         error
		obj         *ResponseStatus
	}{
		// Error responses are only handled if status != http.StatusOK
		{
			name:        "xml",
			status:      http.StatusForbidden,
			contentType: `application/xml; charset="UTF-8"`,
			fixture:     "testdata/fixtures/ResponseStatus.xml",
			err:         ErrResponseNotOk,
			obj: &ResponseStatus{
				XMLVersion:    "1.0",
				XMLNamespace:  "urn:psialliance-org",
				RequestURL:    "/ISAPI/Security/email/parameter",
				StatusCode:    4,
				StatusString:  "Invalid Operation",
				SubStatusCode: "notSupport",
			},
		},
		{
			name:        "json",
			status:      http.StatusForbidden,
			contentType: `application/json; charset="UTF-8"`,
			fixture:     "testdata/fixtures/ResponseStatus.json",
			err:         ErrResponseNotOk,
			obj: &ResponseStatus{
				RequestURL:    "/ISAPI/Security/email/parameter",
				StatusCode:    4,
				StatusString:  "Invalid Operation",
				SubStatusCode: "notSupport",
				ErrorCode:     0,
			},
		},
	}

	// Run test cases
	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Load fixtures
			f, err := os.Open(tc.fixture)
			if err != nil {
				t.Fatalf("error loading test fixtures: %v", err)
			}
			defer f.Close()

			// Read fixture data
			b, err := ioutil.ReadAll(f)
			if err != nil {
				t.Fatalf("error reading file: %v", err)
			}

			// Mock HTTP server
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", tc.contentType)
				w.WriteHeader(tc.status)
				w.Write(b)
			})
			server := httptest.NewServer(handler)

			// Parse URL
			u, err := url.Parse(server.URL)
			if err != nil {
				t.Fatalf("error parsing url: %v", err)
			}

			// Create client object
			c, err := NewClient("", "", "")
			if err != nil {
				t.Fatalf("error creating client: %v", err)
			}

			// Execute request
			b, err = c.Get(u)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected error '%v' but got: %v", tc.err, err)
			}

			// Assert response body
			if b != nil {
				t.Errorf("expected nil body but got: %v", b)
			}
		})
	}
}
