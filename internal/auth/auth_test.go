package auth

import (
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	type test struct {
		input       http.Header
		output      string
		outputError string
	}

	tests := []test{
		{input: http.Header{"Authorization": []string{"ApiKey my-secret-key"}}, output: "my-secret-key", outputError: ""},
		{input: http.Header{}, output: "", outputError: "no authorization header included"},
		{input: http.Header{"Authorization": []string{"ApiKey "}}, output: "", outputError: "malformed authorization header"},
		{input: http.Header{"Authorization": []string{"APIKEY my-secret-key "}}, output: "", outputError: "malformed authorization header"},
	}

	for i, tc := range tests {
		gotOutput, gotError := GetAPIKey(tc.input)
		if gotError != nil {
			if gotError.Error() != tc.outputError {
				t.Errorf("test %d: expected Error: %v, got Error: %v", i+1, tc.outputError, gotError)
				return
			}
		}
		if gotOutput != tc.output {
			t.Errorf("test %d: expected: %v, gotOutput: %v", i+1, tc.output, gotOutput)
		}
	}
}
