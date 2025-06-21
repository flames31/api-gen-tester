package groqclient

import (
	"errors"
	"os"

	"github.com/flames31/groq-client/groq"
	"github.com/joho/godotenv"
)

const MODEL = "llama3-70b-8192"

func Client() (*groq.Client, error) {
	err := godotenv.Load("/Users/rahul31/Desktop/GoProjects/api-gen-tester/internal/groqclient/.env")
	if err != nil {
		return nil, err
	}

	apiKey := os.Getenv("GROQ_APIKEY")

	if apiKey == "" {
		return nil, errors.New("groq api key not set")
	}
	client := groq.NewClient(apiKey, MODEL)

	return client, nil
}
