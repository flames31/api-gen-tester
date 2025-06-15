package parser

import (
	"encoding/json"
	"os"

	"github.com/flames31/api-gen-tester/internal/log"
	types "github.com/flames31/api-gen-tester/internal/types"
)

func ParseJsonFile(file *os.File) (types.ApiTestData, error) {
	var testData types.ApiTestData
	if err := json.NewDecoder(file).Decode(&testData); err != nil {
		log.L().Error("error decoding json file : " + file.Name())
		return types.ApiTestData{}, err
	}

	return testData, nil
}
