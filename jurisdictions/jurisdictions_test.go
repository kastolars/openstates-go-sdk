package jurisdictions

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListJurisdictions(t *testing.T) {
	testCases := []struct {
		description string
		handler     func(w http.ResponseWriter, r *http.Request)
		shouldFail  bool
	}{
		{
			description: "Valid case with both inclusions.",
			handler: func(w http.ResponseWriter, r *http.Request) {
				bytes, _ := json.Marshal(JurisdictionList{})
				io.WriteString(w, string(bytes))
			},
			shouldFail: false,
		},
		{
			description: "Returns nothing",
			handler:     func(w http.ResponseWriter, r *http.Request) {},
			shouldFail:  true,
		},
	}
	for _, tc := range testCases {
		ts := httptest.NewServer(http.HandlerFunc(tc.handler))
		defer ts.Close()

		p := Provider{"", ts.URL}
		_, err := p.ListJurisdictions("state", true, true, 1, 1)
		if (err != nil) != tc.shouldFail {
			t.Errorf("%s failed, error: %v\n", tc.description, err)
		}
	}
}

func TestGetJurisdictionDetails(t *testing.T) {
	testCases := []struct {
		description string
		handler     func(w http.ResponseWriter, r *http.Request)
		shouldFail  bool
	}{
		{
			description: "Valid case",
			handler: func(w http.ResponseWriter, r *http.Request) {
				bytes, _ := json.Marshal(Jurisdiction{})
				io.WriteString(w, string(bytes))
			},
			shouldFail: false,
		},
		{
			description: "Returns nothing",
			handler:     func(w http.ResponseWriter, r *http.Request) {},
			shouldFail:  true,
		},
	}
	for _, tc := range testCases {
		ts := httptest.NewServer(http.HandlerFunc(tc.handler))
		defer ts.Close()

		p := Provider{"", ts.URL}
		_, err := p.GetJurisdictionDetails("", true, true)
		if (err != nil) != tc.shouldFail {
			t.Errorf("%s failed, error: %v\n", tc.description, err)
		}
	}
}
