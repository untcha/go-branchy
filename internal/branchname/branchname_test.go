package branchname

import (
	"errors"
	"strings"
	"testing"
)

func TestNewRequest(t *testing.T) {
	tests := []struct {
		name       string
		issueKey   string
		branchType string
		want       Request
		wantErr    string
	}{
		{
			name:       "normalizes issue key and branch type",
			issueKey:   " abc-1234 ",
			branchType: " Feat ",
			want: Request{
				IssueKey:   "ABC-1234",
				BranchType: "feat",
			},
		},
		{
			name:       "requires issue key",
			branchType: "feat",
			wantErr:    "JIRA issue key is required",
		},
		{
			name:     "requires branch type",
			issueKey: "ABC-1234",
			wantErr:  "branch type is required",
		},
		{
			name:       "rejects slash in branch type",
			issueKey:   "ABC-1234",
			branchType: "feat/api",
			wantErr:    "branch type must not contain slashes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRequest(tt.issueKey, tt.branchType)
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("got error %v, want containing %q", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("got request %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestRequestFromSummary(t *testing.T) {
	request := Request{
		IssueKey:   "ABC-1234",
		BranchType: "feat",
	}

	got, err := request.FromSummary("Add SSO login!")
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}

	want := "feat/ABC-1234/add-sso-login"
	if got != want {
		t.Fatalf("got branch name %q, want %q", got, want)
	}
}

func TestRequestFromSummaryRejectsEmptySlug(t *testing.T) {
	request := Request{
		IssueKey:   "ABC-1234",
		BranchType: "feat",
	}

	_, err := request.FromSummary("!!!")
	if !errors.Is(err, ErrEmptySlug) {
		t.Fatalf("got error %v, want %v", err, ErrEmptySlug)
	}
}

func TestSlugifySummary(t *testing.T) {
	tests := []struct {
		name    string
		summary string
		want    string
		wantErr error
	}{
		{name: "lowercase", summary: "ABC", want: "abc"},
		{name: "special character runs", summary: "this!is@some#text", want: "this-is-some-text"},
		{name: "trims dashes", summary: "-ABC-", want: "abc"},
		{name: "spaces", summary: "Add SSO login", want: "add-sso-login"},
		{name: "empty slug", summary: "!!!", wantErr: ErrEmptySlug},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := slugifySummary(tt.summary)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got error %v, want %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
