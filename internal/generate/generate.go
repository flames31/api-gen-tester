package generate

import (
	"context"
	"fmt"
	"os"

	"github.com/flames31/api-gen-tester/internal/groqclient"
	"github.com/jedib0t/go-pretty/v6/progress"
)

func generateCases(fileName string, genTR *progress.Tracker) (string, error) {
	client, err := groqclient.Client()
	if err != nil {
		return "", fmt.Errorf("error creating new client : %w", err)
	}
	genTR.SetValue(35)

	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("error reading file : %w", err)
	}
	genTR.SetValue(50)
	messages := createPrompt(string(data))
	genTR.SetValue(70)
	resp, err := client.Chat(context.Background(), messages)
	if err != nil {
		return "", fmt.Errorf("error with groq : %w", err)
	}
	genTR.SetValue(90)
	return resp, nil
}
