package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "go-branchy",
	// Version: version,
	Short: "A little helper tool to create git branch names from JIRA tickets - written in Go",
	Long: `A little helper tool to create git branch names from JIRA tickets - written in Go

Branchy is a CLI helper which generates/creates git branch names from JIRA tickets (summary field)
and automatically copies the branch name to the clipboard (e.g. feat/ABC-1234/this-is-my-branch-name)`,

	Example: `go-branchy generate feat ABC-1234
go-branchy g fix ABC-1234`,
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
	rootCmd.AddCommand(newVersionCmd())
}

func initConfig() {
	if err := bindBranchyEnv(); err != nil {
		panic(err)
	}
}

func bindBranchyEnv() error {
	if err := viper.BindEnv("token", "BRANCHY_JIRA_TOKEN"); err != nil {
		return err
	}
	if err := viper.BindEnv("url", "BRANCHY_JIRA_URL"); err != nil {
		return err
	}
	return nil
}
