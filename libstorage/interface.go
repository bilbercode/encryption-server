package libstorage

// Storage provides persistence functionality
type Storage interface {
	// store the data against the provided key
	Store(data []byte, key string) error
	// retrieve the data from persistent storage
	Retrieve(key string) ([]byte, error)
	// Check if the key exists
	Has(key string) (bool, error)
}
