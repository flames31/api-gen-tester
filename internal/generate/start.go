package generate

import (
	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/flames31/api-gen-tester/internal/parser"
	"github.com/flames31/api-gen-tester/internal/tester"
	"go.uber.org/zap"
)

func StartGenerate(fileName string) error {
	log.L().Info("Generating new data for : " + fileName)

	genPW, genTR := startTracker()

	newTestDataStr, err := generateCases(fileName, genTR)
	if err != nil {
		log.L().Error("Error generating new cases : ", zap.Error(err))
		return err
	}

	genTR.SetValue(100)
	genPW.Style().Visibility.Value = false
	genPW.Stop()

	parsedData, err := parser.ParseJsonString(newTestDataStr, fileName)
	if err != nil {
		log.L().Error("error parsing json file : ", zap.Error(err))
		return err
	}

	tester.StartTest(&parsedData)

	if err := parser.WriteJson(&parsedData); err != nil {
		log.L().Error("error writing json to file : ", zap.Error(err))
		return err
	}

	return nil
}
