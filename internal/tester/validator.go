package tester

import (
	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/flames31/api-gen-tester/internal/types"
	"go.uber.org/zap"
)

func startValidator(testCaseChan chan *types.TestCase) {
	log.L().Debug("Starting validator goroutines")
	for i := 0; i < 10; i++ {
		go func(testCaseChan chan *types.TestCase) {
			for tc := range testCaseChan {
				validate(tc)
			}
		}(testCaseChan)
	}
}

func validate(testCase *types.TestCase) {
	log.L().Debug("result :", zap.Int("id", testCase.ID), zap.Int("actual", testCase.Response.StatusCode), zap.Int("expected", testCase.Response.ExpectedStatusCode))
}
