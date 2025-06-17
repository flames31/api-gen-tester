package tester

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/flames31/api-gen-tester/internal/types"
	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
	"go.uber.org/zap"
)

func StartTest(testData *types.ApiTestData) {
	log.L().Debug("entered startTest func")
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(2*time.Second),
		httpclient.WithRetryCount(3),
		httpclient.WithRetrier(heimdall.NewRetrier(heimdall.NewConstantBackoff(500*time.Millisecond, 2*time.Second))),
	)

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10)
	validateChan := make(chan *types.TestCase, 10)

	startValidator(validateChan)

	for i := range testData.TestCases {
		semaphore <- struct{}{}
		wg.Add(1)
		testData.TestCases[i].ID = i + 1
		go sendRequest(testData.BaseURL, &testData.TestCases[i], &wg, semaphore, validateChan, client)
	}
	wg.Wait()
	close(validateChan)
	log.L().Info("Finished processing all test cases.")

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

	url := baseUrl + testCase.Request.Path

	req, err := http.NewRequest(testCase.Request.Method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.L().Error("error creating req :", zap.Error(err))
		return
	}

	for key, val := range testCase.Request.Headers {
		req.Header.Add(key, val)
	}

	res, err := client.Do(req)

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

	testCase.Response.Body = resBody
	testCase.Response.StatusCode = res.StatusCode

	log.L().Debug("Request sent successfully.", zap.Int("id", testCase.ID))

	validateChan <- testCase

	defer res.Body.Close()
}
