package sendgrid

import (
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func sendEmailMessage(message *mail.SGMailV3) error {
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	return err
}
