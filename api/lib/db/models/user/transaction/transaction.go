package transaction

import (
	"time"
)

// ID
type ID string

// ExternalID
type ExternalID string

type ExternalKind string

type Kind string

type Status string
type ExternalStatus string

const (
	KindDebit Kind = "DEBIT"
	KindACH   Kind = "ACH"
)

type Direction string

const (
	DirectionOnramp Direction = "ONRAMP"
	//DirectionAutoLiquidate   Direction = "AUTO_LIQUIDATE"
	//DirectionManualLiquidate Direction = "MANUAL_LIQUIDATE"
)

type Transaction struct {
	ID             ID             `firestore:"id"`
	Kind           Kind           `firestore:"kind"`
	Direction      Direction      `firestore:"direction"`
	Status         Status         `firestore:"status"`
	ExternalID     ExternalID     `firestore:"externalID"`
	ExternalKind   ExternalKind   `firestore:"externalKind"`
	ExternalStatus ExternalStatus `firestore:"status"`
	CreatedAt      time.Time      `firestore:"createdAt"`
	/*
		amount
		currency
		need all da fields (as much as possible which is relevant to user)

	*/
}

type Transactions []Transaction

func (txns Transactions) IDs() []ID {
	var out []ID

	for _, wo := range txns {
		out = append(out, wo.ID)
	}

	return out
}
