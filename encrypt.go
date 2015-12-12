package main

import (
	"github.com/codegangsta/cli"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// Encrypt a file with AES.
var encryptCommand = cli.Command{
	Name:  "encrypt",
	Usage: "encrypt a file",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "input, i",
			Usage: "plaintext file (default: STDIN)",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "encrypted file (default: STDOUT)",
		},
		cli.StringFlag{
			Name:  "key, k",
			Usage: "pre-shared key file",
		},
		cli.BoolFlag{
			Name:  "create, c",
			Usage: "create a new key",
		},
		cli.IntFlag{
			Name:  "size, s",
			Value: 16,
			Usage: "size (in bytes) of the new key (default: 16)",
		},
	},
	Action: func(c *cli.Context) {
		i, err := openInput(c.String("input"))
		if err != nil {
			abortWithError(err)
		}
		defer i.Close()
		o, err := openOutput(c.String("output"))
		if err != nil {
			abortWithError(err)
		}
		defer o.Close()
		var k []byte
		if c.Bool("create") {
			k, err = generateKey(c.String("key"), c.Int("size"))
		} else {
			k, err = openKey(c.String("key"))
		}
		if err != nil {
			abortWithError(err)
		}
		b, err := aes.NewCipher(k)
		if err != nil {
			abortWithError(err)
		}
		iv := make([]byte, aes.BlockSize)
		_, err = rand.Read(iv)
		if err != nil {
			abortWithError(err)
		}
		_, err = o.Write(iv)
		if err != nil {
			abortWithError(err)
		}
		w := &cipher.StreamWriter{
			S: cipher.NewCFBEncrypter(b, iv),
			W: o,
		}
		_, err = io.Copy(w, i)
		if err != nil {
			abortWithError(err)
		}
	},
}
