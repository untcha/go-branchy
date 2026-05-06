package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/untcha/go-branchy/internal/appmeta"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Print the CLI version",
		RunE: func(cmd *cobra.Command, args []string) error {
			printVersion(cmd.OutOrStdout())
			return nil
		},
	}
}

func printVersion(w io.Writer) {
	fmt.Fprintf(w, "Version: %s\n", appmeta.Version)
	fmt.Fprintf(w, "Commit: %s\n", appmeta.Commit)
	fmt.Fprintf(w, "Build date: %s\n", appmeta.BuildDate)
}
