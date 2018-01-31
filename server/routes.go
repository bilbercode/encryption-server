package server

import (
	"net/http"
	"github.com/bilbercode/encryption-server/server/handler"
	"github.com/bilbercode/encryption-server/libstorage"
	"github.com/bilbercode/encryption-server/libcrypto"
	"crypto/rand"
)

// A http route
type Route struct {
	// The name of the route
	Name string
	// The HTTP verb that the route will respond to
	Method string
	// The url path that this route will respond to
	Path string
	// The HTTP handler
	Handler http.HandlerFunc
}

func GetRoutes() ([]Route, error) {

	cryptoService, err := libcrypto.NewYotiCrypto(rand.Reader)
	if err != nil {
		return nil, err
	}
	return []Route{
		{
			Name: "store",
			Path: "/v1/store",
			Method: http.MethodPost,
			Handler: handler.NewStoreHandler(libstorage.NewLocalStorage(), cryptoService).HandleHttp,
		},
		{
			Name: "retrieve",
			Path: "/v1/retrieve",
			Method: http.MethodPost,
			Handler: handler.NewRetrieveHandler(libstorage.NewLocalStorage(), cryptoService).HandleHttp,
		},
		{
			Name: "ping",
			Path: "/",
			Method: http.MethodGet,
			Handler: handler.Ping,
		},
	}, nil
}