package cmd

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestGenerateCommandRejectsExtraArgs(t *testing.T) {
	cmd := newGenerateCmd(func(token, url, issueKey, branchType string) (string, string, error) {
		t.Fatal("generator must not be called for invalid args")
		return "", "", nil
	}, func(text string) error {
		t.Fatal("clipboard must not be called for invalid args")
		return nil
	})
	cmd.SetArgs([]string{"feat", "ABC-1234", "extra"})

	err := cmd.Execute()
	if err == nil || !strings.Contains(err.Error(), "accepts 2 arg(s), received 3") {
		t.Fatalf("got error %v, want exact args error", err)
	}
}

func TestGenerateCommandReportsMissingEnvVars(t *testing.T) {
	t.Setenv("BRANCHY_JIRA_TOKEN", "")
	t.Setenv("BRANCHY_JIRA_URL", "")
	resetViper(t)

	cmd := newGenerateCmd(defaultGenerateBranchNameForTest, func(text string) error {
		t.Fatal("clipboard must not be called when config is invalid")
		return nil
	})
	cmd.SetArgs([]string{"feat", "ABC-1234"})

	err := cmd.Execute()
	if err == nil || !strings.Contains(err.Error(), "BRANCHY_JIRA_TOKEN is required") {
		t.Fatalf("got error %v, want missing token error", err)
	}
}

func TestGenerateCommandPrintsResultAndWritesClipboard(t *testing.T) {
	t.Setenv("BRANCHY_JIRA_TOKEN", "token")
	t.Setenv("BRANCHY_JIRA_URL", "https://jira.example.com")
	resetViper(t)

	var copied string
	cmd := newGenerateCmd(func(token, url, issueKey, branchType string) (string, string, error) {
		if token != "token" {
			t.Fatalf("got token %q", token)
		}
		if url != "https://jira.example.com" {
			t.Fatalf("got url %q", url)
		}
		if issueKey != "ABC-1234" {
			t.Fatalf("got issue key %q", issueKey)
		}
		if branchType != "feat" {
			t.Fatalf("got branch type %q", branchType)
		}
		return "Add SSO login", "feat/ABC-1234/add-sso-login", nil
	}, func(text string) error {
		copied = text
		return nil
	})

	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetArgs([]string{"feat", "ABC-1234"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}

	if copied != "feat/ABC-1234/add-sso-login" {
		t.Fatalf("got copied text %q", copied)
	}

	got := out.String()
	for _, want := range []string{
		"Issue: \t\tABC-1234",
		"Summary: \tAdd SSO login",
		"Branch name: \tfeat/ABC-1234/add-sso-login",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("output %q does not contain %q", got, want)
		}
	}
}

func TestGenerateCommandReturnsClipboardErrorAfterPrintingBranchName(t *testing.T) {
	t.Setenv("BRANCHY_JIRA_TOKEN", "token")
	t.Setenv("BRANCHY_JIRA_URL", "https://jira.example.com")
	resetViper(t)

	wantErr := errors.New("clipboard unavailable")
	cmd := newGenerateCmd(func(token, url, issueKey, branchType string) (string, string, error) {
		return "Add SSO login", "feat/ABC-1234/add-sso-login", nil
	}, func(text string) error {
		return wantErr
	})

	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetArgs([]string{"feat", "ABC-1234"})

	err := cmd.Execute()
	if !errors.Is(err, wantErr) {
		t.Fatalf("got error %v, want wrapped %v", err, wantErr)
	}
	if !strings.Contains(out.String(), "Branch name: \tfeat/ABC-1234/add-sso-login") {
		t.Fatalf("output %q does not include branch name", out.String())
	}
}

func resetViper(t *testing.T) {
	t.Helper()

	viper.Reset()
	t.Cleanup(viper.Reset)
}

func defaultGenerateBranchNameForTest(token, url, issueKey, branchType string) (string, string, error) {
	return generateBranchName(token, url, issueKey, branchType)
}
