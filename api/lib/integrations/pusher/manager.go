package pusher

import (
	"time"

	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/pusher/pusher-http-go"
)

type Manager struct {
	Pusher *pusher.Client
}

type MessageKind string

const (
	MessageKindWyreAccountUpdated        MessageKind = "WYRE_ACCOUNT_UPDATED"
	MessageKindWyrePaymentMethodsUpdated MessageKind = "WYRE_PAYMENT_METHODS_UPDATED"
	MessageKindWyreTransferUpdated       MessageKind = "WYRE_TRANSFER_UPDATED"
	MessageKindProfileStatusUpdated      MessageKind = "PROFILE_STATUS_UPDATED"
)

type Message struct {
	Kind MessageKind `json:"kind"`
	IDs  []string    `json:"ids"`
	At   time.Time   `json:"at"`
}

func (m Manager) Send(userID user.ID, msg *Message) error {
	return m.Pusher.Trigger(string(userID), string(msg.Kind), msg)
}
