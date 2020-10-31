package fortune

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

var tests = []struct {
	name       string
	url        string
	statusCode int
	json       string
	want       string
	errIsNil   bool
}{
	{"Good json response", "https://example.com", 200, `[{"message": "A","id": "5"}]`, "A", true},
	{"Empty json response", "https://example.com", 200, "[]", "", true},
	{"Bad json", "https://example.com", 200, `I am not json`, "", true},
	{"Error 500", "not_a_valid_url", 500, "", "", false},
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

// TestGetFortunes handles table tests
func TestGetFortunes(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewTestClient(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: tt.statusCode,
					Body:       ioutil.NopCloser(bytes.NewBufferString(tt.json)),
					Header:     make(http.Header),
				}
			})

			t.Log("Creating client with url", tt.url)
			api := API{client, tt.url}

			got, err := api.RandomFortune()
			if got != tt.want {
				t.Errorf("RandomFortune() got %v, want %v", got, tt.want)
			}

			t.Log("isErr is", tt.errIsNil, "for", tt.name, "and err is", err)
			if !tt.errIsNil {
				if err == nil {
					t.Errorf("expected error for %s, but err is nil", tt.name)
				}
			}
		})
	}
}
