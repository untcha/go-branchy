package cmd

import (
	"fmt"
	"io"

	"github.com/atotto/clipboard"
	"github.com/untcha/go-branchy/internal/branchy"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

			printGenerateResult(cmd.OutOrStdout(), jiraIssue, summary, branchName)

			if err := writeClipboard(branchName); err != nil {
				return fmt.Errorf("copy branch name to clipboard: %w", err)
			}

			return nil
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(newGenerateCmd(branchy.GenerateBranchName, clipboard.WriteAll))
}

func printGenerateResult(w io.Writer, jiraIssue, summary, branchName string) {
	fmt.Fprintf(w, "Issue: \t\t%s\n", jiraIssue)
	fmt.Fprintf(w, "Summary: \t%s\n", summary)
	fmt.Fprintf(w, "Branch name: \t%s\n", branchName)
}
