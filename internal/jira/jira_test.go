package jira

import (
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		rawURL  string
		wantErr string
	}{
		{
			name:    "creates client",
			token:   "token",
			rawURL:  "https://jira.example.com",
			wantErr: "",
		},
		{
			name:    "requires token",
			rawURL:  "https://jira.example.com",
			wantErr: "BRANCHY_JIRA_TOKEN is required",
		},
		{
			name:    "requires url",
			token:   "token",
			wantErr: "BRANCHY_JIRA_URL is required",
		},
		{
			name:    "requires valid url",
			token:   "token",
			rawURL:  "not-a-url",
			wantErr: "BRANCHY_JIRA_URL must be a valid URL",
		},
		{
			name:    "requires url with host",
			token:   "token",
			rawURL:  "https://",
			wantErr: "BRANCHY_JIRA_URL must include scheme and host",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.token, tt.rawURL)
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("got error %v, want containing %q", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}
			if got == nil {
				t.Fatal("got nil client")
			}
		})
	}
}
