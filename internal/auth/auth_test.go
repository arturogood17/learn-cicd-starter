package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func Test_Auth(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected string
		err      string
	}{
		{
			err: "no authorization header included",
		},
		{
			key:      "Authorization",
			value:    "ApiKey 12345",
			expected: "12345",
			err:      "not expecting an error",
		},
		{
			key:      "Authorization",
			value:    "ApiKey ",
			expected: "",
			err:      "malformed authorization header",
		},
		{
			key:      "Authorization",
			value:    "PepeKey ",
			expected: "",
			err:      "malformed authorization header",
		},
	}

	for index, r := range tests {
		t.Run(fmt.Sprintf("Test #%v", index), func(t *testing.T) {
			h := http.Header{}
			h.Add(r.key, r.value)
			got, f := GetAPIKey(h)
			if got != r.expected {
				t.Fatalf("Expected: %v, got: %v", r.expected, got)
			}
			if f != nil {
				if strings.Contains(f.Error(), r.err) {
					return
				}
				t.Fatalf("Expected: %v, got unexpected error: %v", r.err, f.Error())
				return
			}
		})
	}
}
