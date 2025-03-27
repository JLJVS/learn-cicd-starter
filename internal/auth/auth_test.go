package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		wantKey string
		wantErr string
	}{
		{
			name: "successful api key retrieval",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-api-key"},
			},
			wantKey: "test-api-key",
			wantErr: "",
		},
		{
			name:    "no authorization header",
			headers: http.Header{},
			wantKey: "",
			wantErr: "no authorization header included",
		},
		{
			name: "malformed authorization header - missing ApiKey prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer test-api-key"},
			},
			wantKey: "",
			wantErr: "malformed authorization header",
		},
		{
			name: "malformed authorization header - no space",
			headers: http.Header{
				"Authorization": []string{"ApiKeytest-api-key"},
			},
			wantKey: "",
			wantErr: "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tt.headers)

			// Check error
			if tt.wantErr != "" {
				if err == nil {
					t.Errorf("GetAPIKey() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if err.Error() != tt.wantErr {
					t.Errorf("GetAPIKey() error = %v, wantErr %v", err.Error(), tt.wantErr)
					return
				}
			} else if err != nil {
				t.Errorf("GetAPIKey() unexpected error = %v", err)
				return
			}
			// After checking errors, add this check for the key
			if gotKey != tt.wantKey {
				t.Errorf("GetAPIKey() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}
