package sendemail

import (
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendEmail interface {
	Send(email *mail.SGMailV3) (*rest.Response, error)
}
