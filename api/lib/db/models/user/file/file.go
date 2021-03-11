package file

import "time"

// ID ...
type ID string

// Metadata containes metadata about a cloud bucket file and the DEK used to encrypt/decrypt it
type Metadata struct {
	ID                ID        `firestore:"id"`
	DataEncryptionKey []byte    `firestore:"DEK"`
	MimeType          string    `firestore:"mimeType"`
	Size              int32     `firestore:"size"`
	CreatedAt         time.Time `firestore:"createdAt"`
}

// File represents the returned clear file
type File struct {
	ID        ID
	MimeType  string
	Size      int32
	CreatedAt time.Time
	Body      *[]byte
}
