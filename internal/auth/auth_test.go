package auth_test

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name:    "should give 'no authorization header erro'",
			headers: http.Header{},
			want:    "",
			wantErr: true,
		},
		{
			name:    "should give 'malformed' error when there is no actual key",
			headers: http.Header{"Authorization": []string{"ApiKey"}},
			want:    "",
			wantErr: true,
		},
		{
			name:    "should give 'malformed error when there is no 'ApiKey'",
			headers: http.Header{"Authorization": []string{"Oops Key"}},
			want:    "",
			wantErr: true,
		},
		{
			name:    "should give the actual key",
			headers: http.Header{"Authorization": []string{"ApiKey this_key"}},
			want:    "this_key",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := auth.GetAPIKey(tt.headers)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetAPIKey() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetAPIKey() succeeded unexpectedly")
			}
			if tt.want != got {
				t.Errorf("GetAPIKey()\nGot:\n%v\nWant:\n%v\n", got, tt.want)
			}
		})
	}
}
