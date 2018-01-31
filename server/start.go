package server

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"crypto/tls"
	"crypto/x509"
	"os"
	"errors"
	"io/ioutil"
	"encoding/pem"
	stdLog "log"
)

var (
	errorInvalidCA = errors.New("invalid client CA provided")
)

func Start(addr string, config *Configuration) error {

	routes, err := GetRoutes()
	if err != nil {
		return err
	}

	if config == nil {
		log.Info("starting insecure server")



		router, err := NewRouter(routes)
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("listening on %s", addr)
		// lets keep this on simple as it's only going to be
		// used in dev
		log.Fatal(http.ListenAndServe(addr, router))
	} else {
		log.Info("starting secure server")

		if _, err := os.Stat(config.ClientCA); err != nil {
			log.Fatal(errorInvalidCA)
		}

		caRaw, err := ioutil.ReadFile(config.ClientCA)
		if err != nil {
			log.Fatal(errorInvalidCA)
		}

		block, _ := pem.Decode(caRaw)

		caCertificate, err  := x509.ParseCertificate(block.Bytes)
		if err != nil {
			log.Fatal(errorInvalidCA)
		}

		router, err  := NewRouter(routes)
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("listening on %s", addr)

		caPool := x509.NewCertPool()
		caPool.AddCert(caCertificate)

		w := log.New().Writer()
		// create our http server
		srv := &http.Server{
			Addr: addr,
			Handler: router,
			ErrorLog: stdLog.New(w, "", 0),
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
			TLSConfig: &tls.Config{
				MinVersion:               tls.VersionTLS12,
				CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
				PreferServerCipherSuites: true,
				CipherSuites: []uint16{
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
					tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				},
				ClientCAs: caPool,
				ClientAuth: tls.RequireAndVerifyClientCert,
			},
		}

		log.Fatal(srv.ListenAndServeTLS(config.ServerCertificate, config.ServerCertificateKey))

	}

	return nil
}