package main

import (
	"github.com/codegangsta/cli"

	"crypto/aes"
	"crypto/cipher"
	"io"
)

// Decrypt a file with AES.
var decryptCommand = cli.Command{
	Name:  "decrypt",
	Usage: "decrypt a file",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "input, i",
			Usage: "encrypted file (default: STDIN)",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "plaintext file (default: STDOUT)",
		},
		cli.StringFlag{
			Name:  "key, k",
			Usage: "pre-shared key file",
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
		k, err := openKey(c.String("key"))
		if err != nil {
			abortWithError(err)
		}
		b, err := aes.NewCipher(k)
		if err != nil {
			abortWithError(err)
		}
		iv := make([]byte, aes.BlockSize)
		_, err = i.Read(iv)
		if err != nil {
			abortWithError(err)
		}
		r := &cipher.StreamReader{
			S: cipher.NewCFBDecrypter(b, iv),
			R: i,
		}
		_, err = io.Copy(o, r)
		if err != nil {
			abortWithError(err)
		}
	},
}
