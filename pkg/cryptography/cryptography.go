// inspired by https://github.com/gtank/cryptopasta/blob/master/encrypt.go

// Package cryptography provides symmetric authenticated encryption
// using 256-bit AES-GCM with a random nonce.
package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
)

// encryptionKey used to encrypt and decrypt data.
var encryptionKey = []byte( //nolint:gochecknoglobals
	os.Getenv("ENCRYPTION_KEY"),
)

// Encrypt encrypts data using 256-bit AES-GCM.  This both hides the content of
// the data and provides a check that it hasn't been altered. Output takes the
// form nonce|ciphertext|tag where '|' indicates concatenation.
func Encrypt(plainText []byte, plainKey *string) ([]byte, error) {
	return encrypt(plainText, getKey(plainKey))
}

// Decrypt decrypts data using 256-bit AES-GCM. This both hides the content of
// the data and provides a check that it hasn't been altered. Expects input
// form nonce|ciphertext|tag where '|' indicates concatenation.
func Decrypt(cipherText []byte, plainKey *string) ([]byte, error) {
	return decrypt(cipherText, getKey(plainKey))
}

func encrypt(plainText []byte, key []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("enrypt: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("enrypt: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)

	if err != nil {
		return nil, fmt.Errorf("enrypt: %w", err)
	}

	return gcm.Seal(nonce, nonce, plainText, nil), nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("decrypt: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("decrypt: %w", err)
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("decrypt: malformed ciphertext")
	}

	plainText, err := gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("decrypt: %w", err)
	}

	return plainText, nil
}

func getKey(plainKey *string) []byte {
	var key []byte

	if plainKey == nil {
		key = encryptionKey
	} else {
		key = []byte(*plainKey)
	}

	return key
}
