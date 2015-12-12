## clicrypt

[![MIT License](http://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](http://opensource.org/licenses/MIT)

This simple application simplifies the task of encrypting and decrypting data.

### Usage

The examples below demonstrate basic usage of the application.

#### Encryption

To create an encrypted copy of a file:

    clicrypt encrypt -c -i plain.txt -o cipher.txt

A new pre-shared key which can be used for decrypting the file will be generated and printed to STDERR. Because clicrypt uses STDIN and STDOUT by default, the above command could also be written as:

    clicrypt encrypt -c < plain.txt > cipher.txt

Use the `-k` flag to write the key to disk instead of printing it to STDERR:

    clicrypt encrypt -c -k key ...

To use an existing key instead of generating a new one, omit the `-c` flag.

#### Decryption

To create a decrypted copy of a file:

    clicrypt decrypt -k key -i cipher.txt -o plain.txt

Once again, clicrypt uses STDIN and STDOUT by default, so the above command could also be written as:

    clicrypt decrypt -k key < cipher.txt > plain.txt

#### Using in Pipelines

clicrypt is designed to integrate easily with [pipelines](https://en.wikipedia.org/wiki/Pipeline_(Unix)). For example, to encrypt and compress an entire directory, you could create a pipeline with `tar` and `gz`:

    tar cf - somedir | clicrypt encrypt -c -k key | gzip > cipher.gz

Decompressing and decrypting the directory could then be done with:

    gunzip -c cipher.gz | clicrypt decrypt -k key | tar xf -
