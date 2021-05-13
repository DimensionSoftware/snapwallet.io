package transaction

import (
	"time"
)

// ID
type ID string

// ExternalID
type ExternalID string

type Kind string

const (
	KindDebit Kind = "DEBIT"
	KindACH   Kind = "ACH"
)

type Partner string

const (
	PartnerWyre Partner = "WYRE"
)

type Status string

const (
	StatusQuoted    Status = "QUOTED"
	StatusConfirmed Status = "CONFIRMED"
	StatusCompleted Status = "COMPLETED"
	StatusFailed    Status = "FAILED"
)

type ExternalStatus string

type Direction string

const (
	DirectionOnramp Direction = "ONRAMP"
	//DirectionAutoLiquidate   Direction = "AUTO_LIQUIDATE"
	//DirectionManualLiquidate Direction = "MANUAL_LIQUIDATE"
)

type Transaction struct {
	ID             ID             `firestore:"id"`
	Partner        Partner        `firestore:"partner"`
	Kind           Kind           `firestore:"kind"`
	Direction      Direction      `firestore:"direction"`
	Status         Status         `firestore:"status"`
	ExternalIDs    []ExternalID   `firestore:"externalIDs"`
	ExternalStatus ExternalStatus `firestore:"status"`
	Source         string         `firestore:"source"`         // i.e. "account:AC-WYUR7ZZ6UMU"
	Dest           string         `firestore:"dest"`           // i.e. "bitcoin:14CriXWTRoJmQdBzdikw6tEmSuwxMozWWq"
	SourceName     string         `firestore:"sourceName"`     // i.e. "account:AC-WYUR7ZZ6UMU"
	DestName       string         `firestore:"destName"`       // i.e. "bitcoin:14CriXWTRoJmQdBzdikw6tEmSuwxMozWWq"
	SourceAmount   float64        `firestore:"sourceAmount"`   // i.e. 5
	DestAmount     float64        `firestore:"destAmount"`     // i.e. 0.01
	SourceCurrency string         `firestore:"sourceCurrency"` // i.e. "USD"
	DestCurrency   string         `firestore:"destCurrency"`   // i.e. "BTC"
	Message        string         `firestore:"message"`        // i.e. "Payment for DorianNakamoto@sendwyre.com"
	ExchangeRate   float64        `firestore:"exchangeRate"`   // i.e. 499.00
	TotalFees      float64        `firestore:"totalFees"`
	CreatedAt      time.Time      `firestore:"createdAt"`
	ExpiresAt      time.Time      `firestore:"expiresAt,omitempty"`
	CompletedAt    time.Time      `firestore:"completedAt,omitempty"`
	CancelledAt    time.Time      `firestore:"cancelledAt,omitempty"`
}

type Transactions []Transaction

func (txns Transactions) IDs() []ID {
	var out []ID

	for _, wo := range txns {
		out = append(out, wo.ID)
	}

	return out
}
