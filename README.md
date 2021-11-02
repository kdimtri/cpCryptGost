# Client-server application that uses CryptoPro CSP 5.0 R2 API and GOST encryption algorithm.

## Installing:
First you need to install CryptoPro CSP from https://www.cryptopro.ru/products/csp/downloads
```shell
git clone https://github.com/kdimtri/cpCryptGost
cd cpCryptGost && go build
```
## To run http server:
```shell
./cpCryptGost -s

```
OR
```shell
./cpCryptGost -a localhost:8181 -s
```
## To make a "sign/verify" request to running server:
```shell
./cpCryptGost "message to sign"
```
OR
```shell
./cpCryptGost -a localhost:8181 "message to sign"
```
