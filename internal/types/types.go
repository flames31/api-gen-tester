package types

import "github.com/jedib0t/go-pretty/v6/progress"

type ApiTestData struct {
	BaseURL   string     `json:"base_url"`
	TestCases []TestCase `json:"test_cases"`
}

type TestCase struct {
	ID              int
	Request         Request           `json:"request"`
	Response        Response          `json:"response"`
	ProgressTracker *progress.Tracker `json:"-"`
}

type Request struct {
	Method  string                 `json:"method"`
	Path    string                 `json:"path"`
	Headers map[string]string      `json:"headers"`
	Body    map[string]interface{} `json:"body"`
}

type Response struct {
	StatusCode         int                    `json:"status_code"`
	ExpectedStatusCode int                    `json:"expected_status_code"`
	Body               map[string]interface{} `json:"body"`
}
