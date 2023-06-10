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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.szostok.io/version/extension"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "branchy",
	// Version: version,
	Short: "A little helper tool to create git branch names from JIRA tickets - written in Go",
	Long: `A little helper tool to create git branch names from JIRA tickets - written in Go

Branchy is a CLI helper which generates/creates git branch names from JIRA tickets (summary field)
and automatically copies the branch name to the clipboard (e.g. feat/ABC-1234/this-is-my-branch-name)`,

	Example: `branchy generate feat ABC-1234
branchy g fix ABC-1234`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(extension.NewVersionCobraCmd(extension.WithAliasesOptions("version", "v")))
}

func initConfig() {
	viper.BindEnv("token", "BRANCHY_JIRA_TOKEN")
	viper.BindEnv("url", "BRANCHY_JIRA_URL")
}
