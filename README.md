## clicrypt

[![MIT License](http://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](http://opensource.org/licenses/MIT)

This simple application simplifies the task of encrypting and decrypting data.

### Usage

The examples below demonstrate basic usage of the application.

#### Encryption

To generate a new pre-shared key and encrypt a file with it:

    clicrypt encrypt -c -k key -i plain.txt -o cipher.txt

Because clicrypt uses STDIN and STDOUT by default, the above command could also be written as:

    clicrypt encrypt -c -k key < plain.txt > cipher.txt

To use an existing key instead of generating a new one, omit the `-c` flag.

#### Decryption

To decrypt a file:

    clicrypt decrypt -k key -i cipher.txt -o plain.txt

Once again, clicrypt uses STDIN and STDOUT by default, so the above command could also be written as:

    clicrypt decrypt -k key < cipher.txt > plain.txt

#### Using in Pipelines

clicrypt is designed to integrate easily with [pipelines](https://en.wikipedia.org/wiki/Pipeline_(Unix)). For example, to compress and encrypt an entire directory, you could create a pipeline with `tar`:

    tar czf - somedir | clicrypt encrypt -c -k key > encrypted.tar.gz

Decrypting and decompressing the directory could then be done with:

    clicrypt decrypt -k key < encrypted.tar.gz | tar xzf -
