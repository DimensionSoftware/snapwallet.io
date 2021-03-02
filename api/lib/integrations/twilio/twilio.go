package twilio

import (
	"fmt"
	"log"
	"os"

	"github.com/sfreiberg/gotwilio"
)

const twilioAccountSIDEnvVarName = "TWILIO_ACCOUNT_SID"
const twilioAuthTokenEnvVarName = "TWILIO_AUTH_TOKEN"
const twilioPhoneNumberEnvVarName = "TWILIO_PHONE_NUMBER"

// Config represents the twilio client config
type Config struct {
	AccountSID  string
	AuthToken   string
	PhoneNumber string
}

// ProvideTwilioConfig provides the twilio client config
func ProvideTwilioConfig() (*Config, error) {
	accountSID := os.Getenv(twilioAccountSIDEnvVarName)
	if accountSID == "" {
		return nil, fmt.Errorf("you must set %s", twilioAccountSIDEnvVarName)
	}

	authToken := os.Getenv(twilioAuthTokenEnvVarName)
	if authToken == "" {
		return nil, fmt.Errorf("you must set %s", twilioAuthTokenEnvVarName)
	}

	phoneNumber := os.Getenv(twilioPhoneNumberEnvVarName)
	if phoneNumber == "" {
		return nil, fmt.Errorf("you must set %s", twilioPhoneNumberEnvVarName)
	}

	return &Config{
		AccountSID:  accountSID,
		AuthToken:   authToken,
		PhoneNumber: phoneNumber,
	}, nil
}

// ProvideTwilio provides the twilio client
func ProvideTwilio(config *Config) *gotwilio.Twilio {

	log.Println("🚨 Production Twilio API is activated")
	return gotwilio.NewTwilioClient(config.AccountSID, config.AuthToken)
}
