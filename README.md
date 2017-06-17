
**kb** is a command line tool for saving passwords in keybase.


#### Why
I generally use a password manager, but there are often random secrets I still
have to type into terminals. Available tools I found for storing secrets
use gpg to encrypt/decrypt which is a pain, especially if you use
more than one machine.

This is really easy.

#### How it works

It encrypts passwords with your keybase public key using the `keybase`
command line tool, saving the ciphertext in `/keybase/private/<your username>/credstore`.

#### Installation
First, install and set up [keybase](https://keybase.io/).
```
go get github.com/kingishb/kb
kb init
```

#### Usage
```
$ kb
A key, value store for saving and loading secrets in keybase.

Usage:
  kb COMMAND

Commands:
  init             creates a keystore
  set <key>        save a key
  get <key>        loads value of a key
  generate <key>   auto generates a 12 character random value
  search <string>  lists all keys with substring
  ls               lists all available keys
  rm <key>         removes a key

Example:
  - set key github.com
      kb set github.com

  - get key github.com
      kb get github.com

```
