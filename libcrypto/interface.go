package libcrypto

// Hash output
type Hash []byte

// Crypto provides encryption and decryption functionality
type Crypto interface {
	// encrypt data
	Encrypt(data []byte) (EncryptionResult, error)
	// decrypt data using the key
	Decrypt(data []byte, key []byte) ([]byte, error)
	// Creates a HASH of the id/key composite
	HashIdWithKey(id string, key []byte) Hash
}