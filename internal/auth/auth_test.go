package auth

import("testing"
"net/http")


func TestAuth(t *testing.T){
	tests := []string{
		input http.Header
		want string
		err error
	}{
		{
			input: http.Header{"Authorization": "ApiKey 12345"},
			want: "",
			err: nil,
		},
		{
			input: http.Header{"A": "ApiKey 12345"},
			want: "12345",
			err: errors.New("no authorization header included"),
		},
		{
			input: http.Header{"Authorization": "12345"},
			want: "12345",
			err: errors.New("malformed authorization header"),
		},
	}
}