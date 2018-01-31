package handler

import (
	"net/http"
	"github.com/bilbercode/encryption-server/libstorage"
	"github.com/bilbercode/encryption-server/libcrypto"
)

// Encryption HTTP store handler
type StoreHandler struct {
	storeService libstorage.Storage
	cryptoService libcrypto.Crypto
}

func (h *StoreHandler) HandleHttp(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}


func NewStoreHandler(store libstorage.Storage, crypt libcrypto.Crypto) *StoreHandler {
	return &StoreHandler{
		storeService: store,
		cryptoService: crypt,
	}
}