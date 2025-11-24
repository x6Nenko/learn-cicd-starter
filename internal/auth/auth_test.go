package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError bool
	}{
		{
			name: "valid authorization header",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-secret-key"},
			},
			expectedKey:   "my-secret-key",
			expectedError: false,
		},
		{
			name:          "no authorization header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			if tt.expectedError && err == nil {
				t.Errorf("expected error but got none")
			}

			if !tt.expectedError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if key != tt.expectedKey {
				t.Errorf("expected key %q, got %q", tt.expectedKey, key)
			}
		})
	}
}
