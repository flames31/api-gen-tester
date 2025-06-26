package cmd

import (
	"github.com/flames31/api-gen-tester/internal/generate"
	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	sampleFilePath string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate test cases using sample cases in json file.",
	Long: `The command generates multiple api test cases using LLMs. 
			It uses the cases in sample.json as reference to build new test cases.`,
	Run: runGenerateFunc,
}

func init() {
	initGenerateFlags()
	rootCmd.AddCommand(generateCmd)
}

func runGenerateFunc(cmd *cobra.Command, args []string) {
	log.L().Debug("Entering generate command")
	if sampleFilePath == "" {
		log.L().Fatal("No sample file path provided!!! Shutting down...")
	}

	if err := generate.StartGenerate(sampleFilePath); err != nil {
		log.L().Fatal("Failed to generate testcases!", zap.Error(err))
	}
}

func initGenerateFlags() {
	generateCmd.PersistentFlags().StringVarP(&sampleFilePath, "file", "f", "", "Filepath to sample json")
	generateCmd.MarkFlagRequired("file")
}
