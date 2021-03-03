package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
)

const encryptionKeyKeyEnvVarName = "ENCRYPTION_KEY"

// https://stackoverflow.com/questions/18817336/golang-encrypting-a-string-with-aes-and-base64

// Config is the input config to instantiate the Manager
type Config struct {
	Key string
}

// Manager manages our symmetric at-rest AES-256 encryption
type Manager struct {
	//  AES Cipher Block generated from 32 bytes for AES-256.
	Key cipher.Block
}

// ProvideConfig provides a Config for instantiating the Manager
func ProvideConfig() (*Config, error) {
	key := os.Getenv(encryptionKeyKeyEnvVarName)
	if key == "" {
		return nil, fmt.Errorf("you must set %s", encryptionKeyKeyEnvVarName)
	}

	return &Config{
		Key: key,
	}, nil
}

// NewManager instantiates a new encryption manager
func NewManager(config *Config) (*Manager, error) {
	if len(config.Key) != 32 {
		return nil, fmt.Errorf("encryption key must be 32 bytes in length for AES-256")
	}

	keyBlock, err := aes.NewCipher([]byte(config.Key))
	if err != nil {
		return nil, err
	}

	return &Manager{
		Key: keyBlock,
	}, nil
}

// CipherText is encrypted data
type CipherText = *[]byte

// Encrypt encrypts the cleartext into ciphertext
func (m *Manager) Encrypt(cleartext *[]byte) (CipherText, error) {
	if cleartext == nil {
		return nil, nil
	}

	b64 := base64.StdEncoding.EncodeToString(*cleartext)
	ciphertext := make([]byte, aes.BlockSize+len(b64))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(m.Key, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(*cleartext))

	return &ciphertext, nil

}

// Decrypt decrypts the ciphertext into cleartext
func (m *Manager) Decrypt(ciphertext CipherText) (*[]byte, error) {
	if ciphertext == nil {
		return nil, nil
	}

	c := *ciphertext

	if len(c) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := c[:aes.BlockSize]
	c = c[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(m.Key, iv)

	cfb.XORKeyStream(c, c)

	return &c, nil
}
