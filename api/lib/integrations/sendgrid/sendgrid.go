package sendgrid

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

const sendgridKeyEnvVarName = "SENDGRID_API_KEY"

/*
func sendEmailMessage(message *mail.SGMailV3) error {
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	return err
}

*/
type SendgridApiKey string

func ProvideSendClient(key SendgridApiKey) *sendgrid.Client {
	return sendgrid.NewSendClient(string(key))
}

func ProvideSendClientApiKey() SendgridApiKey {
	sendgridKey := os.Getenv(sendgridKeyEnvVarName)
	if sendgridKey == "" {
		panic(fmt.Sprintf("you must set %s", sendgridKeyEnvVarName))
	}
	return SendgridApiKey(sendgridKey)
}
