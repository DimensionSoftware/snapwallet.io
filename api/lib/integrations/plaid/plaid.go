package plaid

import (
	"fmt"
	"log"
	"os"

	"github.com/plaid/plaid-go/plaid"
)

//client, err := plaid.NewClient(clientOptions)

const plaidClientIDEnvVarName = "PLAID_CLIENT_ID"
const plaidClientSecretEnvVarName = "PLAID_CLIENT_SECRET"

// SendAPIKey represents the secret key for sendgrid
type SendAPIKey string

// ProvideClientOptions provides plaid.ClientOptions
func ProvideClientOptions() (plaid.ClientOptions, error) {
	plaidClientID := os.Getenv(plaidClientIDEnvVarName)
	plaidClientSecret := os.Getenv(plaidClientSecretEnvVarName)

	if plaidClientID == "" {
		return plaid.ClientOptions{}, fmt.Errorf("you must set %s", plaidClientIDEnvVarName)
	}
	if plaidClientSecret == "" {
		return plaid.ClientOptions{}, fmt.Errorf("you must set %s", plaidClientSecretEnvVarName)
	}

	log.Println("🧪 Sandbox Plaid API is activated")

	return plaid.ClientOptions{
		ClientID:    plaidClientID,
		Secret:      plaidClientSecret,
		Environment: plaid.Sandbox,
	}, nil
}
