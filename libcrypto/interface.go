package libcrypto

// Crypto provides encryption and decryption functionality
type Crypto interface {
	// encrypt data
	Encrypt(data []byte) (EncryptionResult, error)
	// decrypt data using the key
	Decrypt(data []byte, key []byte) ([]byte, error)
}