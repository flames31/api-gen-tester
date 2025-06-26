package generate

import (
	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/flames31/groq-client/groq"
)

func createPrompt(testDataStr string) []groq.Message {
	log.L().Debug("creating prompt")
	prompt := []groq.Message{
		{
			Role:    "system",
			Content: "You are a tool that helps expand existing API test cases for automated testing purposes.",
		},
		{
			Role: "user",
			Content: `You are provided with a JSON object that contains a list of test cases, each structured with request and expected response data.
Your job is to:
Analyze the patterns of existing test cases.
Generate 20 additional test cases.
Ensure test cases are relevant variations of existing ones — modify request parameters, use edge cases, or include valid/invalid values.
All the generated test cases should be methods that already exist in the sample json. Example if only "GET" and "POST" methods are present in the sample, all new cases also should only be "GET" and "POST".
Preserve the exact schema and structure of the existing JSON.
Append the new test cases to the existing testCases array without modifying or deleting any existing test cases.
Output the entire JSON file (old + new test cases) in valid JSON format only. Do not include explanation, markdown, or comments.
\n` + testDataStr,
		},
	}

	return prompt
}
