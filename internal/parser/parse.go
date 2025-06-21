package parser

import (
	"encoding/json"
	"os"

	"github.com/flames31/api-gen-tester/internal/log"
	types "github.com/flames31/api-gen-tester/internal/types"
	"go.uber.org/zap"
)

func ParseJsonString(dataStr, fileName string) (types.ApiTestData, error) {
	log.L().Debug("Parsing json file :" + fileName)
	var testData types.ApiTestData
	dataBytes := []byte(dataStr)
	if err := json.Unmarshal(dataBytes, &testData); err != nil {
		log.L().Error("error encoding to json file : "+fileName, zap.Error(err))
		return types.ApiTestData{}, err
	}

	log.L().Debug("Writing new cases to file : results.json")
	err := os.WriteFile("results.json", dataBytes, 0644)
	if err != nil {
		log.L().Error("Error writing to file: results.json", zap.Error(err))
		return types.ApiTestData{}, err
	}
	return testData, nil
}

func ParseJsonBytes(data []byte) (types.ApiTestData, error) {
	var testData types.ApiTestData
	if err := json.Unmarshal(data, &testData); err != nil {
		log.L().Error("error encoding to json : ", zap.Error(err))
		return types.ApiTestData{}, err
	}
	return testData, nil
}
