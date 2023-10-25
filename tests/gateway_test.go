package tests

import (
	"bytes"
	"log"
	"modak-notification-center/services"
	"os"
	"strings"
	"testing"
)

func TestGateway_Send(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	gateway := services.Gateway{}

	user := "example_user"
	message := "Hello, world!"
	gateway.Send(user, message)

	log.SetOutput(os.Stdout)

	expectedLogMessage := "sending message to user example_user\n"
	if !strings.Contains(buf.String(), expectedLogMessage)  {
		t.Errorf("Expected log message to be '%s', but got '%s'", expectedLogMessage, buf.String())
	}
}
