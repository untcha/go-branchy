package cmd

import (
	"fmt"
	"io"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/untcha/go-branchy/internal/branchname"
	"github.com/untcha/go-branchy/internal/jira"
)

type generateBranchNameFunc func(token, url, issueKey, branchType string) (string, string, error)
type clipboardWriteFunc func(text string) error

func newGenerateCmd(generate generateBranchNameFunc, writeClipboard clipboardWriteFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "generate",
		Aliases: []string{"g"},
		Short:   "Generate a branch name from a JIRA issue summary field",
		Long:    `Generate a branch name from a JIRA issue summary field`,
		Example: `go-branchy generate feat ABC-1234`,

		Args: cobra.ExactArgs(2),

		RunE: func(cmd *cobra.Command, args []string) error {
			if err := bindBranchyEnv(); err != nil {
				return err
			}

			jiraToken := viper.GetString("token")
			jiraURL := viper.GetString("url")

			branchType := args[0]
			jiraIssue := args[1]

			summary, branchName, err := generate(
				jiraToken, jiraURL, jiraIssue, branchType,
			)
			if err != nil {
				return err
			}

			if err := printGenerateResult(cmd.OutOrStdout(), jiraIssue, summary, branchName); err != nil {
				return err
			}

			if err := writeClipboard(branchName); err != nil {
				return fmt.Errorf("copy branch name to clipboard: %w", err)
			}

			return nil
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(newGenerateCmd(generateBranchName, clipboard.WriteAll))
}

func printGenerateResult(w io.Writer, jiraIssue, summary, branchName string) error {
	if _, err := fmt.Fprintf(w, "Issue: \t\t%s\n", jiraIssue); err != nil {
		return fmt.Errorf("print issue: %w", err)
	}
	if _, err := fmt.Fprintf(w, "Summary: \t%s\n", summary); err != nil {
		return fmt.Errorf("print summary: %w", err)
	}
	if _, err := fmt.Fprintf(w, "Branch name: \t%s\n", branchName); err != nil {
		return fmt.Errorf("print branch name: %w", err)
	}
	return nil
}

func generateBranchName(token, rawURL, issueKey, branchType string) (string, string, error) {
	request, err := branchname.NewRequest(issueKey, branchType)
	if err != nil {
		return "", "", err
	}

	jiraClient, err := jira.NewClient(token, rawURL)
	if err != nil {
		return "", "", fmt.Errorf("create JIRA client: %w", err)
	}

	summary, err := jiraClient.IssueSummary(request.IssueKey)
	if err != nil {
		return "", "", fmt.Errorf("fetch JIRA issue %s: %w", request.IssueKey, err)
	}

	branchName, err := request.FromSummary(summary)
	if err != nil {
		return "", "", err
	}

	return summary, branchName, nil
}
