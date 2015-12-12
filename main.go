package main

import (
	"github.com/codegangsta/cli"

	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "clicrypt"
	app.Usage = "encrypt and decrypt files with AES"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		encryptCommand,
		decryptCommand,
	}
	app.Run(os.Args)
}
