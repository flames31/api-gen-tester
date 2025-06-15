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
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(2*time.Second),
		httpclient.WithRetryCount(3),
		httpclient.WithRetrier(heimdall.NewRetrier(heimdall.NewConstantBackoff(500*time.Millisecond, 2*time.Second))),
	)

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10)

	for i := range testData.TestCases {
		semaphore <- struct{}{}
		wg.Add(1)
		go sendRequest(testData.BaseURL, &testData.TestCases[i], &wg, semaphore, client, i+1)
	}
	wg.Wait()
	log.L().Info("Finished processing all test cases.")

}

func sendRequest(baseUrl string, testCase *types.TestCase, wg *sync.WaitGroup, semaphore chan struct{}, client *httpclient.Client, id int) {
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

	log.L().Info("Request sent successfully.", zap.Int("id", id))

	defer res.Body.Close()
}
