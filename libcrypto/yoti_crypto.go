package libcrypto

import (
	"io"
	"errors"
	"crypto/aes"
	"sync"
	"crypto/cipher"
	"bytes"
	"golang.org/x/crypto/pbkdf2"
	"crypto/sha512"
	"encoding/hex"
)

var (
	errorBadRand = errors.New("invalid source of rand provided to NewYotiRand")
	errorBadKeyGeneration = errors.New("issue generating correct key length")
	errorBadKeyProvided = errors.New("issue with provided key")
	errorBadCipertext = errors.New("corrupted cipertext provided")
	errorNoEntropy = errors.New("issue reading from the rand pool")
)

// The Yoti crypto
type YotiCrypto struct {
	randSource io.Reader
	randLock *sync.Mutex
}

// get the HEX string of the hash
func (h Hash) ToHexString() string {
	op := make([]byte, 32)
	hex.Encode(op, h)
	return string(op)
}

// Encrypt the data in a consistent manor
func (c *YotiCrypto) Encrypt(data []byte) (EncryptionResult, error) {

	response := EncryptionResult{}
	// lock the rand source to prevent race conditions
	// helpful if provided a rand source which does not do this
	c.randLock.Lock()
	defer c.randLock.Unlock()

	// attempt to generate a secure random crypto key
	key, err := generateNewKey(c.randSource)
	if err != nil {
		return response, errorBadKeyGeneration
	}

	// create a new new cipher block handler
	block, err := aes.NewCipher(key)
	if err != nil {
		return response, errorBadKeyGeneration
	}
	// pad this using a common padding format
	paddedClean := pkcs7Pad(data)

	// make a slice big enough to contain our encrypted data
	ciphertext := make([]byte, aes.BlockSize + len(paddedClean))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(c.randSource, iv); err != nil {
		return response, errorNoEntropy
	}

	cryptoStream := cipher.NewCFBEncrypter(block, iv)
	cryptoStream.XORKeyStream(ciphertext[aes.BlockSize:], paddedClean)

	response.Data = ciphertext
	response.Key = key

	return response, nil

}
// Decrypt the data that was encrypted using YotiCrypto.Encrypt
func (c *YotiCrypto) Decrypt(data []byte, key []byte) ([]byte, error) {

	// check we have been given a valid cipertext to decode
	if (len(data) % aes.BlockSize) != 0 {
		return nil, errorBadCipertext
	}

	if len(data) < aes.BlockSize * 2 {
		// its just to short :(
		return nil, errorBadCipertext
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errorBadKeyProvided
	}

	// grab the IV from the beginning of our stream
	iv := data[:aes.BlockSize]
	// whats left is the cipertext
	dirty := data[aes.BlockSize:]

	cryptoStream := cipher.NewCFBDecrypter(block, iv)
	cryptoStream.XORKeyStream(dirty, dirty)

	clean, err := pkcs7Strip(dirty)
	if err != nil {
		return nil, errorBadCipertext
	}

	return clean, nil
}

// Creates a HASH of the id/key composite
func (c *YotiCrypto) HashIdWithKey(id string, key []byte) Hash {
	// create a clean composite key
	clean := append([]byte(id), key...)
	// just use the lat 8 bits from the end of the AES key
	salt := key[:8]
	/**
	If this was a users credentials we would do this a couple of thousand times
	at least, however it's only used here to completely obfuscate the original
	ID from the one were going to use in storage
	 */
	iterations := int(key[len(key) - 1]) + 22
	// create a hard to derive storage key
	return pbkdf2.Key(clean, salt, iterations, 16, sha512.New)
}



// Generate a new YotiCrypto
func NewYotiCrypto(rand io.Reader) (Crypto, error) {

	b := make([]byte, 1)
	if _, err := rand.Read(b); err != nil {
		return nil, errorBadRand
	}

	return &YotiCrypto{
		randSource: rand,
		randLock: &sync.Mutex{},
	}, nil
}

// Generates a new random key
func generateNewKey(rand io.Reader) ([]byte, error) {
	key := make([]byte, 32)

	if _, err := io.ReadFull(rand, key); err != nil {
		return nil, err
	}

	return key, nil
}

// add PKCS7 padding to the byte slice
func pkcs7Pad(data []byte) []byte {
	// calculate the length of the padding needed
	bit := aes.BlockSize - len(data)%aes.BlockSize
	// pad with the pad length
	padding := bytes.Repeat([]byte{byte(bit)}, bit)
	// return the padded slice
	return append(data, padding...)
}

// Strip PKCS7 padding from the byte slice
func pkcs7Strip(data []byte) ([]byte, error) {
	// get the length of the data to strip
	l := len(data)
	// the last bit will tell us the length it was padded
	iBit := int(data[l- 1])

	if iBit > l {
		// will only occur when encrypted wrong or passed the wrong data to strip of padding
		return nil, errors.New("the pad bit reports longer than the length of the data")
	}
	// return the data < the padding
	return data[:(l - iBit)], nil
}

