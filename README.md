
**kb** is a command line tool for saving passwords in keybase.

![](kb.gif)


#### Why
I use a password manager, but there are often random secrets I still
have to type into terminals. Available tools I found for storing secrets
via command line use gpg to encrypt/decrypt which is a pain.

This is really easy.


#### Installation
First, install and set up [keybase](https://keybase.io/).
```
go get github.com/kingishb/kb
kb init
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
  search <regex>   lists all keys matching regex
  ls               lists all available keys
  rm <key>         removes a key

Example:
  - save key github.com
      kb save github.com

  - get key github.com
      kb get github.com

```
