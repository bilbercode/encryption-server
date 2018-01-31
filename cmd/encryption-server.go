package main

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/bilbercode/encryption-server/command"
	"os"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := cli.NewApp()
	app.Name = "yoti encryption server"
	app.Description = "developer test application"
	app.Version = "1.0.0"
	app.Action = command.Start
	app.Flags = command.StartFlags

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
