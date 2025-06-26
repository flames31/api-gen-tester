package generate

import (
	"fmt"

	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/flames31/api-gen-tester/internal/parser"
	"github.com/flames31/api-gen-tester/internal/tester"
)

func StartGenerate(fileName string) error {
	log.L().Info("Generating new data for : " + fileName)

	genPW, genTR := startTracker()

	newTestDataStr, err := generateCases(fileName, genTR)
	if err != nil {
		return fmt.Errorf("error generating new cases : %w", err)
	}

	genTR.SetValue(100)
	genPW.Style().Visibility.Value = false
	genPW.Stop()

	parsedData, err := parser.ParseJsonString(newTestDataStr, fileName)
	if err != nil {
		return fmt.Errorf("error parsing json file : %w", err)
	}

	tester.StartTest(&parsedData)

	if err := parser.WriteJson(&parsedData); err != nil {
		return fmt.Errorf("error writing json to file : %w", err)
	}

	return nil
}
