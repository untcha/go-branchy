package branchy

import (
	"errors"
	"strings"
	"testing"
)

type fakeIssueSummaryFetcher func(issueKey string) (string, error)

func (f fakeIssueSummaryFetcher) IssueSummary(issueKey string) (string, error) {
	return f(issueKey)
}

func TestGenerateBranchNameWithFetcher(t *testing.T) {
	fetcher := fakeIssueSummaryFetcher(func(issueKey string) (string, error) {
		if issueKey != "ABC-1234" {
			t.Fatalf("got issue key %q", issueKey)
		}
		return "Add SSO login!", nil
	})

	summary, branchName, err := generateBranchName(fetcher, "ABC-1234", "feat")
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}

	if summary != "Add SSO login!" {
		t.Fatalf("got summary %q", summary)
	}
	if branchName != "feat/ABC-1234/add-sso-login" {
		t.Fatalf("got branch name %q", branchName)
	}
}

func TestGenerateBranchNameWithFetcherReturnsFetchError(t *testing.T) {
	wantErr := errors.New("jira unavailable")
	fetcher := fakeIssueSummaryFetcher(func(issueKey string) (string, error) {
		return "", wantErr
	})

	_, _, err := generateBranchName(fetcher, "ABC-1234", "feat")
	if !errors.Is(err, wantErr) {
		t.Fatalf("got error %v, want wrapped %v", err, wantErr)
	}
}

func TestGenerateBranchNameWithFetcherRejectsEmptySlug(t *testing.T) {
	fetcher := fakeIssueSummaryFetcher(func(issueKey string) (string, error) {
		return "!!!", nil
	})

	_, _, err := generateBranchName(fetcher, "ABC-1234", "feat")
	if !errors.Is(err, errEmptySlug) {
		t.Fatalf("got error %v, want %v", err, errEmptySlug)
	}
}

func TestValidateInputs(t *testing.T) {
	tests := []struct {
		name       string
		token      string
		rawURL     string
		issueKey   string
		branchType string
		wantIssue  string
		wantType   string
		wantErr    string
	}{
		{
			name:       "normalizes issue key and branch type",
			token:      "token",
			rawURL:     "https://jira.example.com",
			issueKey:   " abc-1234 ",
			branchType: " Feat ",
			wantIssue:  "ABC-1234",
			wantType:   "feat",
		},
		{
			name:       "requires token",
			rawURL:     "https://jira.example.com",
			issueKey:   "ABC-1234",
			branchType: "feat",
			wantErr:    "BRANCHY_JIRA_TOKEN is required",
		},
		{
			name:       "requires url",
			token:      "token",
			issueKey:   "ABC-1234",
			branchType: "feat",
			wantErr:    "BRANCHY_JIRA_URL is required",
		},
		{
			name:       "requires valid url",
			token:      "token",
			rawURL:     "not-a-url",
			issueKey:   "ABC-1234",
			branchType: "feat",
			wantErr:    "BRANCHY_JIRA_URL must be a valid URL",
		},
		{
			name:       "requires issue key",
			token:      "token",
			rawURL:     "https://jira.example.com",
			branchType: "feat",
			wantErr:    "JIRA issue key is required",
		},
		{
			name:     "requires branch type",
			token:    "token",
			rawURL:   "https://jira.example.com",
			issueKey: "ABC-1234",
			wantErr:  "branch type is required",
		},
		{
			name:       "rejects slash in branch type",
			token:      "token",
			rawURL:     "https://jira.example.com",
			issueKey:   "ABC-1234",
			branchType: "feat/api",
			wantErr:    "branch type must not contain slashes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIssue, gotType, err := validateInputs(tt.token, tt.rawURL, tt.issueKey, tt.branchType)
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("got error %v, want containing %q", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}
			if gotIssue != tt.wantIssue {
				t.Fatalf("got issue %q, want %q", gotIssue, tt.wantIssue)
			}
			if gotType != tt.wantType {
				t.Fatalf("got branch type %q, want %q", gotType, tt.wantType)
			}
		})
	}
}

func TestGetJiraClientReturnsInvalidURLError(t *testing.T) {
	_, err := getJiraClient("token", "://bad")
	if err == nil {
		t.Fatal("got nil error, want invalid JIRA client URL error")
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
		{name: "empty slug", summary: "!!!", wantErr: errEmptySlug},
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

func TestBuildBranchName(t *testing.T) {
	got := buildBranchName("feat", "ABC-1234", "some-text")
	want := "feat/ABC-1234/some-text"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
