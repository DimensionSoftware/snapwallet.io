package file

import "time"

// Metadata containes metadata about a cloud bucket file and the DEK used to encrypt/decrypt it
type Metadata struct {
	ID                string    `firestore:"id"`
	DataEncryptionKey []byte    `firestore:"DEK"`
	MimeType          string    `firestore:"mimeType"`
	Size              int32     `firestore:"size"`
	CreatedAt         time.Time `firestore:"createdAt"`
}
