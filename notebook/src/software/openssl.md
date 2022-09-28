OpenSSL
===

## encrypt/decrypt

Doing some basic encryption
```
openssl enc -aes-256-cbc -md sha512 -pbkdf2 -iter 100000 -salt -in test -out test.enc
```

and decryption
```
openssl enc -aes-256-cbc -d -md sha512 -pbkdf2 -iter 100000 -salt -in test.enc -out test
```
