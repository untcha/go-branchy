/*
Copyright © 2023 Alexander Untch <alexander@untch.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/untcha/go-branchy/internal/branchy"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.design/x/clipboard"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Generate a branch name from a JIRA issue summary field",
	Long:    `Generate a branch name from a JIRA issue summary field`,
	Example: `branchy generate feat ABC-1234`,

	Args: cobra.MinimumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		jiraToken := viper.GetString("token")
		jiraURL := viper.GetString("url")

		// if len(args) < 2 {
		// 	fmt.Println("Please specify a branch type and a JIRA issue.")
		// 	return
		// }

		branchType := args[0]
		jiraIssue := args[1]

		summary, branchName := branchy.GenerateBranchName(
			jiraToken, jiraURL, jiraIssue, branchType,
		)

		err := clipboard.Init()
		if err != nil {
			panic(err)
		}
		clipboard.Write(clipboard.FmtText, []byte(branchName))

		fmt.Printf("Issue: \t\t%s\n", jiraIssue)
		fmt.Printf("Summary: \t%s\n", summary)
		fmt.Printf("Branch name: \t%s\n", branchName)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
