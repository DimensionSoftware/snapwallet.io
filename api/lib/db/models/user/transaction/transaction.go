package transaction

import (
	"strings"
	"time"

	"github.com/khoerling/flux/api/lib/integrations/wyre"
)

// ID
type ID string

// ExternalID
type ExternalID string
type ExternalIDs []ExternalID

func (ids ExternalIDs) Has(targetID ExternalID) bool {
	for _, id := range ids {
		if id == targetID {
			return true
		}
	}
	return false
}

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
	ExternalIDs    ExternalIDs    `firestore:"externalIDs"`
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

func (trx Transaction) EnrichWithWyreTransfer(in wyre.Transfer) Transaction {
	out := trx

	out.Partner = PartnerWyre

	if !out.ExternalIDs.Has(ExternalID(in.ID)) {
		out.ExternalIDs = append(trx.ExternalIDs, ExternalID(in.ID))
	}

	out.ExternalStatus = ExternalStatus(in.Status)
	out.Source = stripWyreObjectPrefix(in.Source)
	out.Dest = stripWyreObjectPrefix(in.Dest)
	out.SourceAmount = in.SourceAmount
	out.DestAmount = in.DestAmount
	out.SourceCurrency = in.SourceCurrency
	out.DestCurrency = in.DestCurrency

	if out.SourceName == "" {
		out.SourceName = in.SourceName
	}

	if out.DestName == "" {
		out.DestName = in.DestName
	}

	out.Message = in.Message

	return out
}

type Transactions []Transaction

func (txns Transactions) IDs() []ID {
	var out []ID

	for _, wo := range txns {
		out = append(out, wo.ID)
	}

	return out
}

func stripWyreObjectPrefix(s string) string {
	parts := strings.Split(s, ":")
	if len(parts) == 0 {
		return ""
	}
	if len(parts) == 1 {
		return parts[0]
	}
	return parts[1]
}
