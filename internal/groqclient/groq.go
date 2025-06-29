package groqclient

import (
	"errors"
	"os"

	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/flames31/groq-client/groq"
	"github.com/joho/godotenv"
)

const MODEL = "llama3-70b-8192"

func Client() (*groq.Client, error) {
	log.L().Debug("creating groq client")
	err := godotenv.Load(".env")
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
