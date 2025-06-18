package tester

import (
	"fmt"
	"sync"

	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/flames31/api-gen-tester/internal/types"
	"go.uber.org/zap"
)

func startValidator(testCaseChan chan *types.TestCase, wg *sync.WaitGroup) {
	log.L().Debug("Starting validator goroutines")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for tc := range testCaseChan {
				log.L().Debug(fmt.Sprintf("Running a validator goroutine for %v", tc.ID))
				tc.ProgressTracker.SetValue(int64(75))
				validate(tc)
			}
		}()
	}
}

func validate(testCase *types.TestCase) {
	log.L().Debug("result :", zap.Int("id", testCase.ID), zap.Int("actual", testCase.Response.StatusCode), zap.Int("expected", testCase.Response.ExpectedStatusCode))
	if testCase.Response.StatusCode == testCase.Response.ExpectedStatusCode {
		testCase.ProgressTracker.MarkAsDone()
	} else {
		testCase.ProgressTracker.MarkAsErrored()
	}
	testCase.ProgressTracker.UpdateMessage(fmt.Sprintf("Request %v : ", testCase.ID))
	testCase.ProgressTracker.SetValue(int64(100))
}
