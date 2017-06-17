
**kb** is a command line tool for saving passwords in keybase.

#### How it works

It encrypts passwords with your keybase public key using the `keybase`
cli, saving the ciphertext in `/keybase/private/<your username>/credstore/<key>`.

#### Installation
First, install and set up [keybase](https://keybase.io/). Then,
```
go get github.com/kingishb/kb
kb init
```

#### Usage
```
$ kb

Usage:
  kb COMMAND

Commands:
  init             creates a keystore
  set <key>        save a key
  get <key>        loads value of a key
  generate <key>   generates & saves 12 character 
                   random value for a key
  search <string>  lists all keys with substring
  ls               lists all available keys
  rm <key>         removes a key

Example:
  - set the key github.com
      kb set github.com

  - get value of the key github.com
      kb get github.com

```
