# hassh-go: SSH server fingerprinting tool

hassh-go is a Golang implementation of the hassh fingerprinting mechanism
developed by Ben Reardon at Salesforce: https://github.com/salesforce/hassh

Unlike the original implementation, hassh-go:
* Does not require a packet capture (uses its own SSH client)
* Has minimal external dependencies (no python, Docker, etc.)
* Compiles to a standalone binary for easy distribution
* Only supports "hasshServer" functionality, i.e. fingerprinting servers (for now)

## Building

To compile `hassh` for the platform you're using, clone this repo and run `make`:
```shell
git clone https://github.com/edutko/hassh-go.git
cd hassh-go
make
```
This will produce a binary in `out/hassh`.

## Usage
```shell
Usage: ./hassh [<options>] <host>[:<port>]
  -v    print kex/cipher/mac details
```
