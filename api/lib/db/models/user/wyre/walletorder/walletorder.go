package walletorder

import (
	"time"
)

// ID
type ID string

// wallet order id; e.g. WO_ELTUVYCAFPG
type WalletOrder struct {
	ID        ID        `firestore:"id"`
	CreatedAt time.Time `firestore:"createdAt"`
}
