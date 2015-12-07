package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
)

// Configuration for the application provided at runtime.
var (
	size     int
	command  string
	filename string
)

// Display usage information for the application.
func printUsage() {
	fmt.Fprintf(os.Stderr, "USAGE\n\thectane [FLAGS] encrypt|decrypt FILE\n\n")
	fmt.Fprintf(os.Stderr, "FLAGS\n")
	flag.PrintDefaults()
}

// Register and parse the configuration flags.
func parseFlags() {
	flag.Usage = printUsage
	flag.IntVar(&size, "size", 16, "PSK size in `bytes`")
	flag.Parse()
	if flag.NArg() != 2 {
		printUsage()
		os.Exit(1)
	}
	command = flag.Arg(0)
	filename = flag.Arg(1)
}

// Encrypt the file.
func encrypt() error {
	psk := make([]byte, size)
	_, err := rand.Read(psk)
	if err != nil {
		return err
	}
	block, err := aes.NewCipher(psk)
	if err != nil {
		return err
	}
	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	if err != nil {
		return err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	inputFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer inputFile.Close()
	outputFilename := filename + ".aes"
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	_, err = outputFile.Write(iv)
	if err != nil {
		return err
	}
	writer := &cipher.StreamWriter{
		S: stream,
		W: outputFile,
	}
	_, err = io.Copy(writer, inputFile)
	if err != nil {
		return err
	}
	fmt.Printf("encrypted \"%s\" with pre-shared key \"%s\"\n",
		outputFilename,
		base64.StdEncoding.EncodeToString(psk))
	return nil
}

// Decrypt the file.
func decrypt() error {
	return nil
}

// Run the application.
func run() error {
	parseFlags()
	switch command {
	case "encrypt":
		return encrypt()
	case "decrypt":
		return decrypt()
	default:
		return fmt.Errorf("unrecognized command \"%s\"", command)
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
