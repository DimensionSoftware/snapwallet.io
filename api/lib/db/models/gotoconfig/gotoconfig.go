package gotoconfig

// ID hashed config, for deduplication
type ID string

// short id for compact / high ec qr codes
type ShortID string

type Config struct {
	ID      ID          `firestore:"id"`
	ShortID ShortID     `firestore:"shortID"`
	Config  interface{} `firestore:"config"`
}
