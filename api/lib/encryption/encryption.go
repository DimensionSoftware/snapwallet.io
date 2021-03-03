package encryption

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/insecurecleartextkeyset"
	"github.com/google/tink/go/keyset"
	"github.com/google/tink/go/tink"
)

const encryptionKeyKeyEnvVarName = "ENCRYPTION_KEY"

// https://stackoverflow.com/questions/18817336/golang-encrypting-a-string-with-aes-and-base64

// Config is the input config to instantiate the Manager
type Config struct {
	// TODO: don't store this in memory and call out to KMS every time (it then acts as our KEK manager)
	MasterKey      string
	AdditionalData []byte
}

// Manager manages our symmetric at-rest encryption
type Manager struct {
	Encryptor      tink.AEAD
	AdditionalData []byte
}

// ProvideConfig provides a Config for instantiating the Manager
func ProvideConfig() (*Config, error) {
	key := os.Getenv(encryptionKeyKeyEnvVarName)
	if key == "" {
		return nil, fmt.Errorf("you must set %s", encryptionKeyKeyEnvVarName)
	}
	return &Config{
		MasterKey: key,
		// TODO: store this in kms later
		AdditionalData: []byte("w8Zp8hAYs1jIkuh2Lc8knbCIN[rWDt.o$=rh'y@agi"),
	}, nil
}

// NewManager instantiates a new encryption manager
func NewManager(config *Config) (*Manager, error) {
	var buf bytes.Buffer
	buf.WriteString(config.MasterKey)

	r := keyset.NewBinaryReader(base64.NewDecoder(base64.RawStdEncoding, &buf))
	// todo: swap out with kms
	key, err := insecurecleartextkeyset.Read(r)
	if err != nil {
		return nil, err
	}

	a, err := aead.New(key)
	if err != nil {
		return nil, err
	}

	return &Manager{
		Encryptor:      a,
		AdditionalData: config.AdditionalData,
	}, nil
}

// CipherText is encrypted data
type CipherText = *[]byte

// Encrypt encrypts the cleartext into ciphertext
func (m *Manager) Encrypt(cleartext *[]byte) (CipherText, error) {
	if cleartext == nil {
		return nil, nil
	}

	ciphertext, err := m.Encryptor.Encrypt(*cleartext, m.AdditionalData)
	if err != nil {
		return nil, err
	}

	return &ciphertext, nil
}

// Decrypt decrypts the ciphertext into cleartext
func (m *Manager) Decrypt(ciphertext CipherText) (*[]byte, error) {
	if ciphertext == nil {
		return nil, nil
	}

	cleartext, err := m.Encryptor.Decrypt(*ciphertext, m.AdditionalData)
	if err != nil {
		return nil, err
	}

	return &cleartext, nil
}
