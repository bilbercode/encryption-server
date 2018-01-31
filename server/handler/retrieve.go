package handler

import (
	"net/http"
	"github.com/bilbercode/encryption-server/libstorage"
	"github.com/bilbercode/encryption-server/libcrypto"
)

// Encryption HTTP retrieve handler
type RetrieveHandler struct {
	storeService libstorage.Storage
	cryptoService libcrypto.Crypto
}

func (h *RetrieveHandler) HandleHttp(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}


func NewRetrieveHandler(store libstorage.Storage, crypt libcrypto.Crypto) *RetrieveHandler {
	return &RetrieveHandler{
		storeService: store,
		cryptoService: crypt,
	}
}


