package tests

import (
	"bytes"
	"modak-notification-center/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func TestSendNotification(t *testing.T) {
	email1 := gofakeit.Email()
	email2 := gofakeit.Email()
	requestBody := []byte(`{"data":{"notifications":[
		{"Type": "news", "Receiver":"`+ email1 +`", "Message": "Hello"},
		{"Type": "status", "Receiver":"`+ email2 +`", "Message": "World"}
	]}}`)
	req, err := http.NewRequest("POST", "/send-notification", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handlers.SendNotification(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Incorrect status code. Expected %d but got %d", http.StatusOK, status)
	}

	expectedResponseBody := `{"data":{"notifications":[{"type":"news","receiver":"`+ email1 +`","message":"Hello","status":"successful"},{"type":"status","receiver":"`+ email2 +`","message":"World","status":"successful"}]}}`
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expectedResponseBody) {
		t.Errorf("Incorrect Response. Expected %s but got %s", expectedResponseBody, rr.Body.String())
	}
	
}

func TestSendNotificationWithoutRequiredField(t *testing.T) {
	email1 := gofakeit.Email()
	requestBody := []byte(`{"data":{"notifications":[
		{"Type": "news", "Receiver":"`+ email1 +`"}
	]}}`)
	req, err := http.NewRequest("POST", "/send-notification", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handlers.SendNotification(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Incorrect status code. Expected %d but got %d", http.StatusBadRequest, status)
	}

	expectedResponseBody := `{"data":{"code":"000.000.001","exception":"required_field_is_missing","status":"error"}}`
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expectedResponseBody) {
		t.Errorf("Incorrect Response. Expected %s but got %s", expectedResponseBody, rr.Body.String())
	}
	
}
