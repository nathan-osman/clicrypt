package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// Configuration for the application provided at runtime.
var (
	key      string
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
	flag.StringVar(&key, "key", "", "`base64`-encoded key used for decryption")
	flag.IntVar(&size, "size", 16, "size in `bytes` of generated PSK")
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
	inputFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer inputFile.Close()
	outputFilename := filename + ".encrypted"
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	psk := make([]byte, size)
	_, err = rand.Read(psk)
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
	_, err = outputFile.Write(iv)
	if err != nil {
		return err
	}
	writer := &cipher.StreamWriter{
		S: cipher.NewCFBEncrypter(block, iv),
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
	inputFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer inputFile.Close()
	outputFilename := strings.TrimSuffix(filename, ".encrypted")
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	if key == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("enter the pre-shared key: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		key = text
	}
	psk, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return err
	}
	block, err := aes.NewCipher(psk)
	if err != nil {
		return err
	}
	iv := make([]byte, aes.BlockSize)
	_, err = inputFile.Read(iv)
	if err != nil {
		return err
	}
	reader := &cipher.StreamReader{
		S: cipher.NewCFBDecrypter(block, iv),
		R: inputFile,
	}
	_, err = io.Copy(outputFile, reader)
	if err != nil {
		return err
	}
	fmt.Printf("decrypted file \"%s\"\n", outputFilename)
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
