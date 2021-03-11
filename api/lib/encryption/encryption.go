package encryption

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
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
	Encryptor tink.AEAD
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
		Encryptor: a,
	}, nil
}

// CipherText is encrypted data
type CipherText = *[]byte

// Encrypt encrypts the cleartext into ciphertext
func (m *Manager) Encrypt(cleartext *[]byte, userID string) (CipherText, error) {
	if cleartext == nil {
		return nil, nil
	}

	ciphertext, err := m.Encryptor.Encrypt(*cleartext, []byte(userID))
	if err != nil {
		return nil, err
	}

	return &ciphertext, nil
}

// Decrypt decrypts the ciphertext into cleartext
func (m *Manager) Decrypt(ciphertext CipherText, userID string) (*[]byte, error) {
	if ciphertext == nil {
		return nil, nil
	}

	cleartext, err := m.Encryptor.Decrypt(*ciphertext, []byte(userID))
	if err != nil {
		return nil, err
	}

	return &cleartext, nil
}

// NewDEK creates new data encryption key
func NewDEK() *keyset.Handle {
	kh, err := keyset.NewHandle(aead.AES256GCMKeyTemplate())
	if err != nil {
		log.Fatal(err)
	}

	return kh
}

// NewEncryptor creates new encryptor
func NewEncryptor(kh *keyset.Handle) tink.AEAD {
	encryptor, err := aead.New(kh)
	if err != nil {
		log.Fatal(err)
	}

	return encryptor
}

// GetEncryptedKeyBytes .
func GetEncryptedKeyBytes(kh *keyset.Handle, masterKey tink.AEAD) *[]byte {
	var buf bytes.Buffer
	err := kh.Write(keyset.NewBinaryWriter(&buf), masterKey)
	if err != nil {
		// should never fail
		log.Fatal(err)
	}

	bytes := buf.Bytes()
	return &bytes
}

// ParseAndDecryptKeyBytes .
func ParseAndDecryptKeyBytes(ciphertext []byte, masterKey tink.AEAD) (*keyset.Handle, error) {
	var buf bytes.Buffer
	_, err := buf.Write(ciphertext)
	if err != nil {
		// should never fail
		log.Fatal(err)
	}

	kh, err := keyset.Read(keyset.NewBinaryReader(&buf), masterKey)
	if err != nil {
		return nil, err
	}

	return kh, nil
}

// DecryptStringIfNonNil ..
func DecryptStringIfNonNil(encryptor tink.AEAD, additionalData []byte, ciphertext *[]byte) (*string, error) {
	if ciphertext == nil {
		return nil, nil
	}

	decrypted, err := DecryptBytesIfNonNil(encryptor, additionalData, ciphertext)
	if err != nil {
		return nil, err
	}

	s := string(*decrypted)

	return &s, nil
}

// DecryptBytesIfNonNil ..
func DecryptBytesIfNonNil(encryptor tink.AEAD, additionalData []byte, ciphertext *[]byte) (*[]byte, error) {
	if ciphertext == nil {
		return nil, nil
	}

	decrypted, err := encryptor.Decrypt(*ciphertext, additionalData)
	if err != nil {
		return nil, err
	}

	return &decrypted, nil
}

// EncryptStringIfNonNil ..
func EncryptStringIfNonNil(encryptor tink.AEAD, additionalData []byte, cleartext *string) (*[]byte, error) {
	if cleartext == nil {
		return nil, nil
	}

	b := []byte(*cleartext)

	encrypted, err := EncryptBytesIfNonNil(encryptor, additionalData, &b)
	if err != nil {
		return nil, err
	}

	return encrypted, nil
}

// EncryptBytesIfNonNil ..
func EncryptBytesIfNonNil(encryptor tink.AEAD, additionalData []byte, cleartext *[]byte) (*[]byte, error) {
	if cleartext == nil {
		return nil, nil
	}

	encrypted, err := encryptor.Encrypt(*cleartext, additionalData)
	if err != nil {
		return nil, err
	}

	return &encrypted, nil
}
