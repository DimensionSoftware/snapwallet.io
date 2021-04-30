package paymentmethod

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user/plaid/item"
)

// ID ...
type ID string

// PaymentMethod ...
type PaymentMethod struct {
	ID                    ID             `firestore:"id"`
	PlaidItemID           item.ID        `firestore:"plaidItemID,omitempty"`
	PlaidAccountID        item.AccountID `firestore:"plaidAccountID,omitempty"`
	Status                string         `firestore:"status"`
	Name                  string         `firestore:"name"`
	Last4                 string         `firestore:"last4"`
	ChargeableCurrencies  []string       `firestore:"chargeableCurrencies"`
	DepositableCurrencies []string       `firestore:"depositableCurrencies"`
	CreatedAt             time.Time      `firestore:"createdAt"`
	UpdatedAt             time.Time      `firestore:"updatedAt"`
}
