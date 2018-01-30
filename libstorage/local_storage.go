package libstorage

import (
	"sync"
	"errors"
	"fmt"
)

type LocalStorage struct {
	items map[string][]byte
	lock sync.RWMutex
}

func (s *LocalStorage) Store(data []byte, key string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items[key] = data
	return nil
}

func (s *LocalStorage) Retrieve(key string) ([]byte, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	value, ok := s.items[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("key `%s` does not exist", key))
	}

	return value, nil
}


func (s *LocalStorage) Has(key string) (bool, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if _, ok := s.items[key]; ok {
		return true, nil
	}
	return false, nil
}

// provides a new LocalStorage
func NewLocalStorage() Storage {
	return &LocalStorage{
		items: map[string][]byte{},
		lock: sync.RWMutex{},
	}
}