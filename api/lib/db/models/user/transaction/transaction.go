package transaction

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/integrations/wyre"
	proto "github.com/khoerling/flux/api/lib/protocol"
	"github.com/lithammer/shortuuid"
)

// EncryptedTransaction is the at-rest form of transactions
type EncryptedTransaction struct {
	ID                ID          `firestore:"id"`
	ExternalIDs       ExternalIDs `firestore:"externalIDs"`
	Partner           Partner     `firestore:"partner"`
	Kind              Kind        `firestore:"kind"`
	Direction         Direction   `firestore:"direction"`
	Status            Status      `firestore:"status"`
	DataEncryptionKey []byte      `firestore:"DEK"`
	DataEncrypted     []byte      `firestore:"dataEncrypted"`
	CreatedAt         time.Time   `firestore:"createdAt"`
}

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
	ID             ID                 `json:"id"`
	Partner        Partner            `json:"partner"`
	Kind           Kind               `json:"kind"`
	Direction      Direction          `json:"direction"`
	Status         Status             `json:"status"`
	ExternalIDs    ExternalIDs        `json:"externalIDs"`
	ExternalStatus ExternalStatus     `json:"externalStatus"`
	Source         string             `json:"source"`         // i.e. "account:AC-WYUR7ZZ6UMU"
	Dest           string             `json:"dest"`           // i.e. "bitcoin:14CriXWTRoJmQdBzdikw6tEmSuwxMozWWq"
	SourceName     string             `json:"sourceName"`     // i.e. "account:AC-WYUR7ZZ6UMU"
	DestName       string             `json:"destName"`       // i.e. "bitcoin:14CriXWTRoJmQdBzdikw6tEmSuwxMozWWq"
	SourceAmount   float64            `json:"sourceAmount"`   // i.e. 5
	DestAmount     float64            `json:"destAmount"`     // i.e. 0.01
	SourceCurrency string             `json:"sourceCurrency"` // i.e. "USD"
	DestCurrency   string             `json:"destCurrency"`   // i.e. "BTC"
	Message        string             `json:"message"`        // i.e. "Payment for DorianNakamoto@sendwyre.com"
	ExchangeRate   float64            `json:"exchangeRate"`   // i.e. 499.00
	TotalFees      float64            `json:"totalFees"`
	Fees           map[string]float64 `json:"fees"`
	CreatedAt      time.Time          `json:"createdAt"`
	ExpiresAt      time.Time          `json:"expiresAt,omitempty"`
	CompletedAt    time.Time          `json:"completedAt,omitempty"`
	CancelledAt    time.Time          `json:"cancelledAt,omitempty"`
}

// WithDefaults provides defaults for User
func (trx Transaction) WithDefaults() Transaction {
	newTRX := trx

	if trx.ID == "" {
		newTRX.ID = ID(shortuuid.New())
	}

	if (trx.CreatedAt == time.Time{}) {
		newTRX.CreatedAt = time.Now()
	}

	return newTRX
}

func (trx Transaction) EnrichWithWyreTransfer(in *wyre.Transfer) Transaction {
	out := trx

	out.Partner = PartnerWyre
	// todo: infer from input status
	out.Status = StatusConfirmed

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
	out.CreatedAt = fromEpochMS(in.CreatedAt)
	// todo: is this right? is closed at same as completed at?
	//out.CompletedAt = fromEpochMS(in.ClosedAt)

	if out.SourceName == "" {
		out.SourceName = in.SourceName
	}

	if out.DestName == "" {
		out.DestName = in.DestName
	}

	out.Message = in.Message

	return out
}

func (trx Transaction) EnrichWithWyreTransferDetail(in *wyre.TransferDetail) Transaction {
	out := trx

	out.Partner = PartnerWyre
	// todo: infer from input status
	out.Status = StatusConfirmed

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
	out.Message = in.Message
	out.ExchangeRate = in.ExchangeRate
	out.TotalFees = in.TotalFees
	out.Fees = in.Fees
	out.CreatedAt = fromEpochMS(in.CreatedAt)
	out.ExpiresAt = fromEpochMS(in.ExpiresAt)
	out.CompletedAt = fromEpochMS(in.CompletedAt)
	out.CancelledAt = fromEpochMS(in.CancelledAt)

	if in.DestCurrency == "BTC" || in.DestCurrency == "ETH" {
		// doing this because a better name may have already been set
		if out.DestName == "" {
			out.DestName = fmt.Sprintf("%s Address: %s", in.DestCurrency, out.Dest)
		}
	}

	// todo: we can do better than this by looking up payment methods
	if in.SourceCurrency == "USD" {
		if out.SourceName == "" {

			// doing this because a better name may have already been set
			if out.Kind == KindACH {
				out.SourceName = fmt.Sprintf("Bank Account: %s", out.Source)
			}

			if out.Kind == KindDebit {
				out.SourceName = fmt.Sprintf("Debit Card: %s", out.Source)
			}
		}
	}

	return out
}

func (trx Transaction) EnrichWithCreateWalletOrderReservationResponse(in *wyre.CreateWalletOrderReservationResponse) Transaction {
	out := trx

	out.Partner = PartnerWyre
	// this input object always indicates a quoted status
	out.Status = StatusQuoted

	if !out.ExternalIDs.Has(ExternalID(in.Reservation)) {
		out.ExternalIDs = append(trx.ExternalIDs, ExternalID(in.Reservation))
	}

	return out
}

func (trx Transaction) EnrichWithCreateWalletOrderRequest(in *wyre.CreateWalletOrderRequest) Transaction {
	out := trx

	out.Partner = PartnerWyre

	{
		num := in.DebitCard.Number
		last4 := num[len(num)-3:]
		out.SourceName = fmt.Sprintf("Debit Card ending with %s", last4)
	}

	return out
}

func (trx Transaction) EnrichWithWalletOrderReservation(in *wyre.WalletOrderReservation) Transaction {
	out := trx

	out.Partner = PartnerWyre

	if out.DestAmount == 0 {
		out.DestAmount = in.DestAmount
	}
	out.Fees = in.Quote.Fees

	total := 0.0
	for _, fee := range out.Fees {
		total += fee
	}
	out.TotalFees = total

	return out
}

func (trx Transaction) EnrichWithWalletOrder(in *wyre.WalletOrder) Transaction {
	out := trx

	out.Partner = PartnerWyre
	// todo: infer from input status
	out.Status = StatusConfirmed

	if !out.ExternalIDs.Has(ExternalID(in.ID)) {
		out.ExternalIDs = append(trx.ExternalIDs, ExternalID(in.ID))
	}

	out.ExternalStatus = ExternalStatus(in.Status)
	// todo: figure out what source makes sense here as a generated fill in
	//out.Source = ""
	out.Dest = stripWyreObjectPrefix(in.Dest)
	out.SourceAmount = in.SourceAmount
	out.SourceCurrency = in.SourceCurrency
	out.DestCurrency = in.DestCurrency

	if out.SourceName == "" {
		out.SourceName = fmt.Sprintf("Debit Card (Wallet Order): %s", in.ID)
	}

	if out.DestName == "" {
		if in.DestCurrency == "BTC" || in.DestCurrency == "ETH" {
			// doing this because a better name may have already been set
			if out.DestName == "" {
				out.DestName = fmt.Sprintf("%s Address: %s", in.DestCurrency, out.Dest)
			}
		}
	}

	return out
}

func (trx Transaction) AsProto() *proto.Transaction {
	return &proto.Transaction{
		Id:             string(trx.ID),
		Partner:        string(trx.Partner),
		Kind:           string(trx.Kind),
		Direction:      string(trx.Direction),
		Status:         string(trx.Status),
		Source:         trx.Source,
		Dest:           trx.Dest,
		SourceName:     trx.SourceName,
		DestName:       trx.DestName,
		SourceAmount:   trx.SourceAmount,
		DestAmount:     trx.DestAmount,
		SourceCurrency: trx.SourceCurrency,
		DestCurrency:   trx.DestCurrency,
		Message:        trx.Message,
		ExchangeRate:   trx.ExchangeRate,
		TotalFees:      trx.TotalFees,
		CreatedAt:      formatAsRFC3339(trx.CreatedAt),
		ExpiresAt:      formatAsRFC3339(trx.ExpiresAt),
		CompletedAt:    formatAsRFC3339(trx.CompletedAt),
		CancelledAt:    formatAsRFC3339(trx.CancelledAt),
	}
}

type Transactions []Transaction

func (txns Transactions) AsProto() *proto.Transactions {
	var out []*proto.Transaction

	for _, txn := range txns {
		out = append(out, txn.AsProto())
	}

	return &proto.Transactions{
		Transactions: out,
	}
}

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

func fromEpochMS(epochMS int64) time.Time {
	if epochMS == 0 {
		return time.Time{}
	}

	return time.Unix(epochMS/1000, 0)
}

// Encrypt encrypts the transaction
func (transaction Transaction) Encrypt(m *encryption.Manager, userID user.ID) (*EncryptedTransaction, error) {
	dekH := encryption.NewDEK()
	dek := encryption.NewEncryptor(dekH)

	jsonData, err := json.Marshal(&transaction)
	if err != nil {
		return nil, err
	}

	encryptedData, err := dek.Encrypt(jsonData, []byte(userID))
	if err != nil {
		return nil, err
	}

	return &EncryptedTransaction{
		ID:                transaction.ID,
		ExternalIDs:       transaction.ExternalIDs,
		Partner:           transaction.Partner,
		Kind:              transaction.Kind,
		Direction:         transaction.Direction,
		Status:            transaction.Status,
		DataEncryptionKey: *encryption.GetEncryptedKeyBytes(dekH, m.Encryptor),
		DataEncrypted:     encryptedData,
		CreatedAt:         transaction.CreatedAt,
	}, nil
}

// Decrypt decrypts the transaction
func (enc EncryptedTransaction) Decrypt(m *encryption.Manager, userID user.ID) (*Transaction, error) {
	dekH, err := encryption.ParseAndDecryptKeyBytes(enc.DataEncryptionKey, m.Encryptor)
	if err != nil {
		return nil, err
	}
	dek := encryption.NewEncryptor(dekH)

	jsonData, err := dek.Decrypt(enc.DataEncrypted, []byte(userID))
	if err != nil {
		return nil, err
	}

	var transaction Transaction
	err = json.Unmarshal(jsonData, &transaction)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func formatAsRFC3339(t time.Time) string {
	if (t == time.Time{}) {
		return ""
	}

	return t.Format(time.RFC3339)
}
