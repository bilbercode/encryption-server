package server

// Configuration for the server
type Configuration struct {
	// The location of the servers x509 certificate
	ServerCertificate string
	// The location of the servers private key
	ServerCertificateKey string
	// The client CA certificate location to validate the client
	ClientCA string
}
