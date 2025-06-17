package generate

import (
	"fmt"
	"os"

	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/flames31/api-gen-tester/internal/parser"
	"github.com/flames31/api-gen-tester/internal/tester"
)

func Generate(fileName string) error {
	log.L().Info("Generating data for : " + fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.L().Error("error opening file : " + err.Error())
	}
	defer file.Close()
	parsedData, err := parser.ParseJsonFile(file)
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.L().Debug("Calling tester.StartTest")
	tester.StartTest(&parsedData)

	return nil
}
