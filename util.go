package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

// Display an error and terminate the application.
func abortWithError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}

// Open the input file. If filename is empty, return STDIN instead.
func openInput(filename string) (*os.File, error) {
	if filename == "" {
		return os.Stdin, nil
	}
	return os.Open(filename)
}

// Open the output file. If filename is empty, return STDOUT instead.
func openOutput(filename string) (*os.File, error) {
	if filename == "" {
		return os.Stdout, nil
	}
	return os.Create(filename)
}

// Generate a new pre-shared key of the requested size.
func generateKey(filename string, size int) ([]byte, error) {
	var (
		w   *os.File
		err error
	)
	if filename == "" {
		w = os.Stderr
	} else {
		w, err = os.Create(filename)
		if err != nil {
			return nil, err
		}
		defer w.Close()
	}
	b := make([]byte, size)
	_, err = rand.Read(b)
	if err != nil {
		return nil, err
	}
	if filename == "" {
		fmt.Fprintf(os.Stderr, "Key generated: %s\n",
			base64.StdEncoding.EncodeToString(b))
	} else {
		w := base64.NewEncoder(base64.StdEncoding, w)
		defer w.Close()
		_, err = w.Write(b)
		if err != nil {
			return nil, err
		}
	}
	return b, nil
}

// Open the pre-shared key.
func openKey(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, f))
	if err != nil {
		return nil, err
	}
	return b, nil
}
