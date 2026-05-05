package branchy

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/andygrunwald/go-jira"
)

const jiraClientTimeout = 15 * time.Second

var errEmptySlug = errors.New("issue summary does not contain any branch-safe characters")

var nonBranchSafeChars = regexp.MustCompile(`[^a-zA-Z0-9]+`)

type issueSummaryFetcher interface {
	IssueSummary(issueKey string) (string, error)
}

type jiraIssueSummaryFetcher struct {
	client *jira.Client
}

func (f jiraIssueSummaryFetcher) IssueSummary(issueKey string) (string, error) {
	return getJiraIssueSummary(f.client, issueKey)
}

func GenerateBranchName(token, rawURL, issueKey, branchType string) (string, string, error) {
	issueKey, branchType, err := validateInputs(token, rawURL, issueKey, branchType)
	if err != nil {
		return "", "", err
	}

	c, err := getJiraClient(token, rawURL)
	if err != nil {
		return "", "", fmt.Errorf("create JIRA client: %w", err)
	}

	return generateBranchName(jiraIssueSummaryFetcher{client: c}, issueKey, branchType)
}

func generateBranchName(fetcher issueSummaryFetcher, issueKey, branchType string) (string, string, error) {
	summary, err := fetcher.IssueSummary(issueKey)
	if err != nil {
		return "", "", fmt.Errorf("fetch JIRA issue %s: %w", issueKey, err)
	}

	slug, err := slugifySummary(summary)
	if err != nil {
		return "", "", err
	}

	return summary, buildBranchName(branchType, issueKey, slug), nil
}

func validateInputs(token, rawURL, issueKey, branchType string) (string, string, error) {
	if strings.TrimSpace(token) == "" {
		return "", "", errors.New("BRANCHY_JIRA_TOKEN is required")
	}
	if strings.TrimSpace(rawURL) == "" {
		return "", "", errors.New("BRANCHY_JIRA_URL is required")
	}
	if err := validateJiraURL(rawURL); err != nil {
		return "", "", err
	}

	issueKey = strings.ToUpper(strings.TrimSpace(issueKey))
	if issueKey == "" {
		return "", "", errors.New("JIRA issue key is required")
	}

	branchType = strings.ToLower(strings.TrimSpace(branchType))
	if branchType == "" {
		return "", "", errors.New("branch type is required")
	}
	if strings.Contains(branchType, "/") {
		return "", "", errors.New("branch type must not contain slashes")
	}

	return issueKey, branchType, nil
}

func validateJiraURL(rawURL string) error {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return fmt.Errorf("BRANCHY_JIRA_URL must be a valid URL: %w", err)
	}
	if u.Scheme == "" || u.Host == "" {
		return errors.New("BRANCHY_JIRA_URL must include scheme and host")
	}
	return nil
}

func getJiraClient(token, rawURL string) (*jira.Client, error) {
	tp := jira.BearerAuthTransport{
		Token: token,
	}
	httpClient := tp.Client()
	httpClient.Timeout = jiraClientTimeout

	jiraClient, err := jira.NewClient(httpClient, rawURL)
	if err != nil {
		return nil, err
	}

	return jiraClient, nil
}

func getJiraIssueSummary(jiraClient *jira.Client, issueKey string) (string, error) {
	issue, _, err := jiraClient.Issue.Get(issueKey, nil)
	if err != nil {
		return "", err
	}
	return issue.Fields.Summary, nil
}

func slugifySummary(summary string) (string, error) {
	s := toLowerCase(summary)
	s = replaceAllSpecialCharacters(s)
	s = trimDash(s)
	if s == "" {
		return "", errEmptySlug
	}
	return s, nil
}

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func replaceAllSpecialCharacters(s string) string {
	return nonBranchSafeChars.ReplaceAllString(s, "-")
}

func trimDash(s string) string {
	return strings.Trim(s, "-")
}

func buildBranchName(branchType, issueKey, s string) string {
	return fmt.Sprintf("%s/%s/%s", branchType, issueKey, s)
}
