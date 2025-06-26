package tester

import (
	"fmt"
	"sync"

	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/flames31/api-gen-tester/internal/tracker"
	"github.com/flames31/api-gen-tester/internal/types"
	"github.com/gojek/heimdall/v7/httpclient"
	"go.uber.org/zap"
)

func StartTest(testData *types.ApiTestData) {
	log.L().Debug("entered startTest func")
	client := newHeimdallClient()
	pw := tracker.NewTracker()

	var sendWG sync.WaitGroup
	var valWG sync.WaitGroup
	semaphore := make(chan struct{}, 10)
	validateChan := make(chan *types.TestCase, 10)

	startValidator(validateChan, &valWG)

	go pw.Render()

	for i := range testData.TestCases {
		semaphore <- struct{}{}
		sendWG.Add(1)
		updateTestCase(testData, i, pw)

		go sendRequest(testData.BaseURL, &testData.TestCases[i], &sendWG, semaphore, validateChan, client)
	}

	sendWG.Wait()
	close(validateChan)
	valWG.Wait()
	log.L().Info("Finished processing all test cases.")
	pw.Stop()
}

func sendRequest(baseUrl string, testCase *types.TestCase, wg *sync.WaitGroup, semaphore chan struct{}, validateChan chan *types.TestCase, client *httpclient.Client) {
	log.L().Debug("starting go routine to send req", zap.Int("id", testCase.ID), zap.Any("test_case", testCase))
	defer wg.Done()
	defer func() { <-semaphore }()

	if err := send(testCase, baseUrl, client); err != nil {
		log.L().Error(fmt.Sprintf("error with req %v", testCase.ID), zap.Error(err))
	}

	log.L().Debug("Request sent successfully.", zap.Int("id", testCase.ID))

	testCase.ProgressTracker.SetValue(int64(50))

	validateChan <- testCase
}
