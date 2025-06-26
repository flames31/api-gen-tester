package tester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/flames31/api-gen-tester/internal/types"
	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/jedib0t/go-pretty/v6/progress"
)

func newHeimdallClient() *httpclient.Client {
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(2*time.Second),
		httpclient.WithRetryCount(3),
		httpclient.WithRetrier(heimdall.NewRetrier(heimdall.NewConstantBackoff(500*time.Millisecond, 2*time.Second))),
	)

	return client
}

func updateTestCase(testData *types.ApiTestData, i int, pw progress.Writer) {
	testData.TestCases[i].ID = i + 1
	testData.TestCases[i].ProgressTracker = &progress.Tracker{
		Message: fmt.Sprintf("Request %v", i+1),
		Total:   100,
		Units:   progress.UnitsDefault,
	}
	pw.AppendTracker(testData.TestCases[i].ProgressTracker)
}

func send(testCase *types.TestCase, baseUrl string, client *httpclient.Client) error {
	log.L().Debug("entered send func")
	reqBody, err := json.Marshal(testCase.Request.Body)
	if err != nil {
		return fmt.Errorf("error marshalling req body : %w", err)

	}

	testCase.ProgressTracker.SetValue(10)

	url := baseUrl + testCase.Request.Path

	req, err := http.NewRequest(testCase.Request.Method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("error creating req : %w", err)
	}

	testCase.ProgressTracker.SetValue(int64(20))

	for key, val := range testCase.Request.Headers {
		req.Header.Add(key, val)
	}

	res, err := client.Do(req)

	testCase.ProgressTracker.SetValue(int64(30))

	if err != nil {
		return fmt.Errorf("error sending req : %w", err)
	}

	defer res.Body.Close()
	var resBody map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		return fmt.Errorf("error unmarshalling res body : %w", err)
	}

	testCase.ProgressTracker.SetValue(int64(40))

	testCase.Response.Body = resBody
	testCase.Response.StatusCode = res.StatusCode

	return nil
}
