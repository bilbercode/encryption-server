package libcrypto

import (
	"io"
)

// The Yoti crypto
type YotiCrypto struct {
	randSource io.Reader
}

// Encrypt the data in a consistent manor
func (c *YotiCrypto) Encrypt(data []byte) (EncryptionResult, error) {
	panic("not implemented")
}
// Decrypt the data that was encrypted using YotiCrypto.Encrypt
func (c *YotiCrypto) Decrypt(data []byte, key []byte) ([]byte, error) {
	panic("not implemented")
}

// Generate a new YotiCrypto
func NewYotiCrypto(rand io.Reader) (Crypto, error) {
	panic("not implemented")
}