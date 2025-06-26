package generate

import (
	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/flames31/groq-client/groq"
)

func createPrompt(testDataStr string) []groq.Message {
	log.L().Debug("creating tracker")
	prompt := []groq.Message{
		{
			Role:    "system",
			Content: "You are an assistant that generates and appends API test cases to the provided JSON without deleting any information.",
		},
		{
			Role:    "user",
			Content: "Generate 10 test cases for the endpoints as seen below. Do remove any data already present. Only add new test cases. Just give me the json string, nothing else.\n" + testDataStr,
		},
	}

	return prompt
}
