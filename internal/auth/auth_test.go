package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestAuth(t *testing.T) {
	type test struct {
		input       http.Header
		output      string
		outputError error
	}

	tests := []test{
		{input: http.Header{"Authorization": []string{"ApiKey my-secret-key"}}, output: "my-secret-key", outputError: nil},
		{input: http.Header{}, output: "", outputError: errors.New("no authorization header included")},
		{input: http.Header{"Authorization": []string{"ApiKey "}}, output: "", outputError: errors.New("malformed authorization header")},
		{input: http.Header{"Authorization": []string{"APIKEY my-secret-key "}}, output: "", outputError: errors.New("malformed authorization header")},
	}

	for i, tc := range tests {
		gotOutput, gotError := GetAPIKey(tc.input)
		if !reflect.DeepEqual(tc.output, gotOutput) {
			t.Fatalf("test %d: expected: %v, gotOutput: %v, gotError: %v", i+1, tc.output, gotOutput, gotError)
		}
	}
}
