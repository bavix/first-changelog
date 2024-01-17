package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/bavix/first-changelog/internal/app"
)

var rootCmd = &cobra.Command{
	Use:   "first-changelog",
	Short: "Changelog generated automatically",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			fmt.Println(app.GenChangelog(cmd.Context(), arg))
		}
	},
}

func Execute(ctx context.Context) {
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}
