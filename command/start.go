package command

import (
	"github.com/bilbercode/encryption-server/server"
	"gopkg.in/urfave/cli.v1"
)

var StartFlags = []cli.Flag{
	cli.StringFlag{
		Name: "addr, lister-addr",
		Usage: "Network address for the server to listen on",
		Value: "0.0.0.0:443",
		EnvVar: "SERVER_ADDR",
	},
	cli.StringFlag{
		Name:   "cert, certificate",
		Usage:  "Location of the server certificate",
		Value:  "",
		EnvVar: "SERVER_CERTIFICATE",
	},
	cli.StringFlag{
		Name:   "key, certificate-key",
		Usage:  "Location of the server certificate key",
		Value:  "",
		EnvVar: "SERVER_CERTIFICATE_KEY",
	},
	cli.StringFlag{
		Name:   "ca",
		Usage:  "Location of the CA",
		Value:  "",
		EnvVar: "SERVER_CA",
	},
}

func Start(c *cli.Context) error {

	config := convertCliFlagsToConfiguration(c)

	return server.Start(c.String("addr"), config)
}

func convertCliFlagsToConfiguration(c *cli.Context) *server.Configuration {
	cert := c.String("cert")
	key := c.String("key")
	ca := c.String("ca")

	if cert == "" || key == "" || ca == "" {
		return nil
	}

	return &server.Configuration{
		ServerCertificate:    cert,
		ServerCertificateKey: key,
		ClientCA:             ca,
	}
}
