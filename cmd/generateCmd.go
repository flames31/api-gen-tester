package cmd

import (
	"os"

	"github.com/flames31/api-gen-tester/internal/generate"
	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/spf13/cobra"
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
		log.L().Error("No sample file path provided")
		os.Exit(1)
	}

	if err := generate.StartGenerate(sampleFilePath); err != nil {
		log.L().Error("Failed to generate testcases")
		os.Exit(1)
	}
}

func initGenerateFlags() {
	generateCmd.PersistentFlags().StringVarP(&sampleFilePath, "file", "f", "", "Filepath to sample json")
	generateCmd.MarkFlagRequired("file")
}
