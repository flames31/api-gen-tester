package tester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/flames31/api-gen-tester/internal/types"
	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/jedib0t/go-pretty/v6/progress"
	"go.uber.org/zap"
)

func StartTest(testData *types.ApiTestData) {
	log.L().Debug("entered startTest func")
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(2*time.Second),
		httpclient.WithRetryCount(3),
		httpclient.WithRetrier(heimdall.NewRetrier(heimdall.NewConstantBackoff(500*time.Millisecond, 2*time.Second))),
	)

	pw := progress.NewWriter()
	pw.SetTrackerLength(25)
	pw.SetUpdateFrequency(100 * time.Millisecond)
	pw.SetStyle(progress.StyleDefault)
	pw.SetAutoStop(false)

	var sendWG sync.WaitGroup
	var valWG sync.WaitGroup
	semaphore := make(chan struct{}, 10)
	validateChan := make(chan *types.TestCase, 10)

	startValidator(validateChan, &valWG)

	go pw.Render()

	for i := range testData.TestCases {
		semaphore <- struct{}{}
		sendWG.Add(1)
		testData.TestCases[i].ID = i + 1
		testData.TestCases[i].ProgressTracker = &progress.Tracker{
			Message: fmt.Sprintf("Request %v", i+1),
			Total:   100,
			Units:   progress.UnitsDefault,
		}
		pw.AppendTracker(testData.TestCases[i].ProgressTracker)
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

	reqBody, err := json.Marshal(testCase.Request.Body)
	if err != nil {
		log.L().Error("error marshalling req body :", zap.Error(err))
		return
	}

	testCase.ProgressTracker.SetValue(int64(10))

	url := baseUrl + testCase.Request.Path

	req, err := http.NewRequest(testCase.Request.Method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.L().Error("error creating req :", zap.Error(err))
		return
	}

	testCase.ProgressTracker.SetValue(int64(20))

	for key, val := range testCase.Request.Headers {
		req.Header.Add(key, val)
	}

	res, err := client.Do(req)

	testCase.ProgressTracker.SetValue(int64(30))

	if err != nil {
		log.L().Error("error sending req :", zap.Error(err))
		return
	}

	var resBody map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		log.L().Error("error unmarshalling res body :", zap.Error(err))
		return
	}

	testCase.ProgressTracker.SetValue(int64(40))

	testCase.Response.Body = resBody
	testCase.Response.StatusCode = res.StatusCode

	log.L().Debug("Request sent successfully.", zap.Int("id", testCase.ID))

	testCase.ProgressTracker.SetValue(int64(50))

	validateChan <- testCase

	defer res.Body.Close()
}
