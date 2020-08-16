package hikvision

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestClient_GetDeviceInfoCapabilities(t *testing.T) {
	// Create test cases
	cases := []struct {
		name     string
		encoding string
		fixture  string
	}{
		{"default", "xml", "./testdata/fixtures/DeviceInfoCap.xml"},
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

			// Read test fixture
			b, err := ioutil.ReadAll(f)
			if err != nil {
				t.Fatalf("error reading file: %v", err)
			}

			// Derive content type
			var contentType string
			switch tc.encoding {
			case "xml":
				contentType = `application/xml; charset="UTF-8"`
			case "json":
				contentType = `application/json; charset="UTF-8"`
			default:
				t.Fatalf("error deriving content type, expect either of 'xml' or 'json' but got: %v", tc.encoding)
			}

			// Mock HTTP server
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", contentType)
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			})
			server := httptest.NewServer(handler)

			// Execute request
			c, err := NewClient("localhost", "", "")
			if err != nil {
				t.Fatalf("error creating client: %v", err)
			}
			c.BaseURL = server.URL
			got, err := c.GetDeviceInfoCapabilities()
			if err != nil {
				t.Fatalf("error executing request: %v", err)
			}

			// Unmarshal test fixture
			var want *DeviceInfoCap
			switch tc.encoding {
			case "xml":
				if err := xml.Unmarshal(b, &want); err != nil {
					t.Fatalf("error unmarshalling fixture: %v", err)
				}
			case "json":
				if err := json.Unmarshal(b, &want); err != nil {
					t.Fatalf("error unmarshalling fixture: %v", err)
				}
			default:
				t.Fatalf("error deriving content type, expect either of 'xml' or 'json' but got: %v", tc.encoding)
			}

			// Assert struct
			if !reflect.DeepEqual(got, want) {
				t.Errorf("error struct not equal: got %+v want %+v", got, want)
			}
		})
	}
}
