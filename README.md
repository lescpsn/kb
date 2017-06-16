
**kb** is a command line tool for saving passwords in keybase.


#### Why
Although I use a password manager, there are random secrets I still
have to type into terminals. Available tools I found for storing secrets
via command line use gpg to encrypt/decrypt which is a pain.


#### Installation
First, install and set up [keybase](https://keybase.io/).
```
go get github.com/kingishb/kb
```

#### Usage
```
$ kb

An key, value store for saving and loading secrets in keybase.

Usage:
  kb COMMAND

Commands:
  init             creates a keystore
  save <key>       save a key
  get <key>        loads value of a key
  generate <key>   auto generates a 12 character random value
  list             lists all available keys
  search <regex>   lists all keys matching regex

Example:
  - save key github.com
      kb save github.com

  - get key github.com
      kb get github.com
```
