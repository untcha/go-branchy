package branchy

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/andygrunwald/go-jira"
)

func GenerateBranchName(token, url, issueKey, branchType string) (string, string) {
	c, _ := getJiraClient(token, url)
	summary, _ := getJiraIssueSummary(c, issueKey)

	s := toLowerCase(summary)
	s = replaceAllSpecialCharacters(s)
	s = trimDash(s)
	branchName := buildBranchName(branchType, issueKey, s)

	return summary, branchName
}

func getJiraClient(token, url string) (*jira.Client, error) {
	tp := jira.BearerAuthTransport{
		Token: token,
	}

	jiraClient, err := jira.NewClient(tp.Client(), url)
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

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func replaceAllSpecialCharacters(s string) string {
	return regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "-")
}

func trimDash(s string) string {
	return strings.Trim(s, "-")
}

func buildBranchName(branchType, issueKey, s string) string {
	return fmt.Sprintf("%s/%s/%s", branchType, issueKey, s)
}
