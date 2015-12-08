## clicrypt

[![MIT License](http://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](http://opensource.org/licenses/MIT)

This simple application simplifies the task of encrypting and decrypting a single file.

### Usage

Encrypting a file is as simple as:

    $ clicrypt encrypt somefile.txt
    encrypted "somefile.txt.encrypted" with pre-shared key "..."

Be sure to hang on to the pre-shared key since it will be needed to decrypt the file later:

    $ clicrypt decrypt somefile.txt.encrypted
    enter the pre-shared key: [...]
    decrypted file "somefile.txt"

It is possible to pass the pre-shared key as an argument to `clicrypt`, but doing so [may be a security vulnerability](http://unix.stackexchange.com/q/8223/1049).
