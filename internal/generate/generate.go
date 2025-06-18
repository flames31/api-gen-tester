package generate

import (
	"os"

	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/flames31/api-gen-tester/internal/parser"
	"github.com/flames31/api-gen-tester/internal/tester"
	"go.uber.org/zap"
)

func Generate(fileName string) error {
	log.L().Info("Generating data for : " + fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.L().Error("error opening file : ", zap.Error(err))
		return err
	}

	defer file.Close()
	parsedData, err := parser.ParseJsonFile(file)
	if err != nil {
		log.L().Error("error parsing json file : ", zap.Error(err))
		return err
	}
	log.L().Debug("Calling tester.StartTest")
	tester.StartTest(&parsedData)

	if err := parser.WriteJson(&parsedData); err != nil {
		log.L().Error("error writing json to file : ", zap.Error(err))
		return err
	}

	return nil
}
