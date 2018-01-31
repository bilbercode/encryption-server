package intergration

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"crypto/tls"
	"os"
	"path"
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
)

var _ = Describe("Server", func() {


	var client *http.Client

	BeforeSuite(func() {

		cwd, _ := os.Getwd()
		certRoot := path.Join(cwd, "..", ".cfssl")
		cCert, err := tls.LoadX509KeyPair(path.Join(certRoot, "client.pem"), path.Join(certRoot, "client-key.pem"))
		if err != nil {
			panic(err)
		}

		caRaw, err := ioutil.ReadFile(path.Join(certRoot, "ca.pem"))
		if err != nil {
			panic(err)
		}

		block, _ := pem.Decode(caRaw)

		caCertificate, err  := x509.ParseCertificate(block.Bytes)
		if err != nil {
			panic(err)
		}

		pool := x509.NewCertPool()
		pool.AddCert(caCertificate)

		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					MinVersion:               tls.VersionTLS12,
					CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
					PreferServerCipherSuites: true,
					CipherSuites: []uint16{
						tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
						tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
						tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
						tls.TLS_RSA_WITH_AES_256_CBC_SHA,
					},
					Certificates: []tls.Certificate{ cCert },
					RootCAs: pool,
				},
			},
		}
	})

	Describe("Authorization", func() {
		It("should respond to an authorised http ping", func() {
			request, err:= http.NewRequest(http.MethodGet, "https://0.0.0.0/", nil)
			Expect(err).ToNot(HaveOccurred())
			res, err := client.Do(request)
			Expect(err).ToNot(HaveOccurred())
			Expect(res.StatusCode).To(Equal(200))
		})
	})
})
