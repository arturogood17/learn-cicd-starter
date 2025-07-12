package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	tests := []struct {
		key   string
		value string
		want  string
		err   string
	}{
		{
			key:   "Authorization",
			value: "ApiKey 12345",
			want:  "12345",
			err:   "",
		},
		{
			key:   "A",
			value: "ApiKey 12345",
			want:  "",
			err:   "no authorization header included",
		},
		{
			key:   "Authorization",
			value: "12345",
			want:  "",
			err:   "malformed authorization header",
		},
	}

	for i, test := range tests {
		h := http.Header{}
		t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
			h[test.key] = []string{test.value}
			got, err := GetAPIKey(h)
			if err != nil {
				if err.Error() != test.err {
					t.Errorf("unexpected error")
				}
			}
			if got != test.want {
				t.Fatalf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
