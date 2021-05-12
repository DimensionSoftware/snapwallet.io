package walletorderreservation

import (
	"time"
)

// ID
type ID string

// wallet order reservation id
type WalletOrderReservation struct {
	ID                       ID        `firestore:"id"`
	WalletOrderReservationID string    `firestore:"walletOrderReservationID,omitempty"`
	CreatedAt                time.Time `firestore:"createdAt"`
	UpdatedAt                time.Time `firestore:"updatedAt"`
}
