package widgetconfig

import "github.com/khoerling/flux/api/lib/protocol"

// ID hashed config, for deduplication
type ID string

// short id for compact / high ec qr codes
type ShortID string

// Metadata containes metadata about a cloud bucket file and the DEK used to encrypt/decrypt it
type Config struct {
	ID      ID                         `firestore:"id"`
	ShortID ShortID                    `firestore:"shortId"`
	Config  *protocol.SnapWidgetConfig `firestore:"config"`
}
