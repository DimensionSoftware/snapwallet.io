package paymentmethod

import (
	"time"
)

// ID ...
type ID string

// PaymentMethod ...
type PaymentMethod struct {
	ID                    ID        `firestore:"id"`
	PlaidItemID           string    `firestore:"plaidItemID,omitempty"`
	PlaidAccountID        string    `firestore:"plaidAccountID,omitempty"`
	Status                string    `firestore:"status"`
	Name                  string    `firestore:"name"`
	Last4                 string    `firestore:"last4"`
	ChargeableCurrencies  []string  `firestore:"chargeableCurrencies"`
	DepositableCurrencies []string  `firestore:"depositableCurrencies"`
	CreatedAt             time.Time `firestore:"createdAt"`
	UpdatedAt             time.Time `firestore:"updatedAt"`
}
