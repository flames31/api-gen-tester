package parser

import (
	"encoding/json"
	"os"

	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/flames31/api-gen-tester/internal/types"
)

func WriteJson(resultData *types.ApiTestData) error {
	log.L().Debug("Writing output to results.json")
	file, err := os.Create("results.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(resultData)
}
