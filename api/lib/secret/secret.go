package secret

import "github.com/google/tink/go/tink"

// Encrypted ...
type Encrypted interface {
	Decrypt(tink.AEAD, []byte) (*interface{}, error)
}

// Clear ...
type Clear interface {
	Encrypt(tink.AEAD, *interface{}) ([]byte, error)
}
