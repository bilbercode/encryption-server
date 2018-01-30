package libstorage

import "sync"

type LocalStorage struct {
	items map[string][]byte
	lock sync.RWMutex
}

func (s *LocalStorage) Store(data []byte, key string) error {
	panic("not implemented")
}

func (s *LocalStorage) Retrieve(key string) ([]byte, error) {
	panic("not implemented")
}


func (s *LocalStorage) Has(key string) (bool, error) {
	panic("not implemented")
}

func NewLocalStorage() Storage {
	panic("not implemented")
}