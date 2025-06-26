package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/flames31/api-gen-tester/internal/log"
	types "github.com/flames31/api-gen-tester/internal/types"
)

func ParseJsonString(dataStr, fileName string) (types.ApiTestData, error) {
	log.L().Debug("Parsing json file :" + fileName)
	var testData types.ApiTestData
	dataBytes := []byte(dataStr)
	if err := json.Unmarshal(dataBytes, &testData); err != nil {
		return types.ApiTestData{}, fmt.Errorf("error encoding to json file : %w", err)
	}

	log.L().Debug("Writing new cases to file : results.json")
	err := os.WriteFile("results.json", dataBytes, 0644)
	if err != nil {
		return types.ApiTestData{}, fmt.Errorf("error writing to file: %w", err)
	}
	return testData, nil
}

func ParseJsonBytes(data []byte) (types.ApiTestData, error) {
	var testData types.ApiTestData
	if err := json.Unmarshal(data, &testData); err != nil {
		return types.ApiTestData{}, fmt.Errorf("error encoding to json : %w", err)
	}
	return testData, nil
}
