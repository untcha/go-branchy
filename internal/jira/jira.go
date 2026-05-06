package jira

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	gojira "github.com/andygrunwald/go-jira"
)

const clientTimeout = 15 * time.Second

type Client struct {
	client *gojira.Client
}

func NewClient(token, rawURL string) (*Client, error) {
	if strings.TrimSpace(token) == "" {
		return nil, errors.New("JIRA_TOKEN is required")
	}
	if strings.TrimSpace(rawURL) == "" {
		return nil, errors.New("JIRA_HOST is required")
	}
	if err := validateURL(rawURL); err != nil {
		return nil, err
	}

	tp := gojira.BearerAuthTransport{
		Token: token,
	}
	httpClient := tp.Client()
	httpClient.Timeout = clientTimeout

	jiraClient, err := gojira.NewClient(httpClient, rawURL)
	if err != nil {
		return nil, err
	}

	return &Client{client: jiraClient}, nil
}

func (c *Client) IssueSummary(issueKey string) (string, error) {
	issue, _, err := c.client.Issue.Get(issueKey, nil)
	if err != nil {
		return "", err
	}
	return issue.Fields.Summary, nil
}

func validateURL(rawURL string) error {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return fmt.Errorf("JIRA_HOST must be a valid URL: %w", err)
	}
	if u.Scheme == "" || u.Host == "" {
		return errors.New("JIRA_HOST must include scheme and host")
	}
	return nil
}
