package middleware

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := json.Marshal(readAndRestoreBody(r))
		log.Printf("Request: %s %s %s, Body: %s", r.Method, r.URL, r.RemoteAddr, string(body))

		lw := &responseLogger{ResponseWriter: w}

		next.ServeHTTP(lw, r)

		log.Printf("Response: Status %d, Size %d bytes", lw.status, lw.size)
	})
}

type responseLogger struct {
	http.ResponseWriter
	status int
	size   int
}

func (rl *responseLogger) WriteHeader(statusCode int) {
	rl.status = statusCode
	rl.ResponseWriter.WriteHeader(statusCode)
}

func (rl *responseLogger) Write(b []byte) (int, error) {
	n, err := rl.ResponseWriter.Write(b)
	rl.size += n
	return n, err
}

func readAndRestoreBody(r *http.Request) interface{} {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body request: %v", err)
	}

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("Error analyzing body request: %v", err)
	}

	r.Body = ioutil.NopCloser(strings.NewReader(string(body)))
	return result
}