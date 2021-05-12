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

type WalletOrders []WalletOrder

func (wos WalletOrders) IDs() []string {
	var out []string

	for _, wo := range wos {
		out = append(out, string(wo.ID))
	}

	return out
}
