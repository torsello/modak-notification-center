package tests

import (
	"bytes"
	"log"
	"modak-notification-center/middleware"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	originalLogOutput := log.Writer()
	defer func() {
		log.SetOutput(originalLogOutput)
	}()
	
	var logOutput bytes.Buffer
	log.SetOutput(&logOutput)

	defer func() {
		log.SetOutput(originalLogOutput)
	}()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Response Body"))
	})

	req, err := http.NewRequest("POST", "/test", strings.NewReader(`{"key": "value"}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	middleware.LoggingMiddleware(handler).ServeHTTP(rr, req)

	expectedRequestLog := "Request: POST /test , Body: {\"key\":\"value\"}\n"
	expectedResponseLog := "Response: Status 200, Size 13 bytes\n"

	actualLogOutput := logOutput.String()

	if !strings.Contains(actualLogOutput, expectedRequestLog) || !strings.Contains(actualLogOutput, expectedResponseLog) {
		t.Errorf("Incorrect middleware registration. Expected:\n%s\nObteined:\n%s", expectedRequestLog+expectedResponseLog, actualLogOutput)
	}
	
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Incorrect status code. Expected %d, obtained %d", http.StatusOK, status)
	}

	expectedResponseBody := "Response Body"
	if body := rr.Body.String(); body != expectedResponseBody {
		t.Errorf("Wrong response body. Expected %s, obtained %s", expectedResponseBody, body)
	}
}
