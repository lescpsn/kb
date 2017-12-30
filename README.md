
a small  command line password manager.

#### How it works

It encrypts passwords with your keybase public key using the [keybase
cli](https://keybase.io/docs/command_line), saving the ciphertext in `$HOME/.kb/<key>`.

#### Installation
Install [keybase](https://keybase.io/) and [go](https://golang.org/doc/install). Then,
```
go get github.com/kingishb/kb
kb init
```
#### Usage
```
$ kb --help
Usage:
  kb COMMAND

Commands:
  init             creates a keystore
  set <key>        save a key
  get <key>        loads value of a key
  generate <key>   generates & saves 20 character
                   random value for a key
  search <string>  lists all keys with substring
  ls               lists all available keys
  rm <key>         removes a key
```
