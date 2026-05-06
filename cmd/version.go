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
			return printVersion(cmd.OutOrStdout())
		},
	}
}

func printVersion(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "Version: %s\n", appmeta.Version); err != nil {
		return fmt.Errorf("print version: %w", err)
	}
	if _, err := fmt.Fprintf(w, "Commit: %s\n", appmeta.Commit); err != nil {
		return fmt.Errorf("print commit: %w", err)
	}
	if _, err := fmt.Fprintf(w, "Build date: %s\n", appmeta.BuildDate); err != nil {
		return fmt.Errorf("print build date: %w", err)
	}
	return nil
}
