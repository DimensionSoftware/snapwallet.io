package sendgrid

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

const sendgridKeyEnvVarName = "SENDGRID_API_KEY"

// SendAPIKey represents the secret key for sendgrid
type SendAPIKey string

// ProvideSendClient returns a sendgrid client
func ProvideSendClient(key SendAPIKey) *sendgrid.Client {
	return sendgrid.NewSendClient(string(key))
}

// ProvideSendClientAPIKey returns a sendgrid api key string
func ProvideSendClientAPIKey() (SendAPIKey, error) {
	sendgridKey := os.Getenv(sendgridKeyEnvVarName)
	if sendgridKey == "" {
		return "", fmt.Errorf("you must set %s", sendgridKeyEnvVarName)
	}

	log.Println("🚨 Production SendGrid API is activated")

	return SendAPIKey(sendgridKey), nil
}
