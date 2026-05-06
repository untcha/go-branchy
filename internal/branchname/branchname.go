package branchname

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var ErrEmptySlug = errors.New("issue summary does not contain any branch-safe characters")

var nonBranchSafeChars = regexp.MustCompile(`[^a-zA-Z0-9]+`)

type Request struct {
	IssueKey   string
	BranchType string
}

func NewRequest(issueKey, branchType string) (Request, error) {
	issueKey = strings.ToUpper(strings.TrimSpace(issueKey))
	if issueKey == "" {
		return Request{}, errors.New("JIRA issue key is required")
	}

	branchType = strings.ToLower(strings.TrimSpace(branchType))
	if branchType == "" {
		return Request{}, errors.New("branch type is required")
	}
	if strings.Contains(branchType, "/") {
		return Request{}, errors.New("branch type must not contain slashes")
	}

	return Request{
		IssueKey:   issueKey,
		BranchType: branchType,
	}, nil
}

func (r Request) FromSummary(summary string) (string, error) {
	slug, err := slugifySummary(summary)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s/%s", r.BranchType, r.IssueKey, slug), nil
}

func slugifySummary(summary string) (string, error) {
	s := strings.ToLower(summary)
	s = nonBranchSafeChars.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	if s == "" {
		return "", ErrEmptySlug
	}
	return s, nil
}
