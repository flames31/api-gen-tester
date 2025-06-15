package cmd

import (
	"os"

	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/spf13/cobra"
)

var (
	logLevel string
)

var rootCmd = &cobra.Command{
	Use:   "api-gen-tester",
	Short: "A cli tool to test API with AI generated cases.",
	Long: `This tool takes in a few API sample cases and uses AI to generate additional edge cases.
	The sample test cases can be provided using json.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return log.Init(logLevel)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&logLevel, "loglevel", "info", "Log level (debug, info, warn, error)")
}
