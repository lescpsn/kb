
kb is a command line password manager using the keybase filesystem for
password storage. Use it for passwords you have to type into your
terminal, or just as an encrypted key value store.

#### Installation
First, install and set up [keybase](https://keybase.io/).
```
go get github.com/kingishb/kb
```

#### Usage
```
$ kb

An encrypted key, value store for saving and loading secrets.

Usage:
  kb COMMAND

Commands:
  init             creates a keystore
  save <key>       save a key
  get <key>        loads value of a key
  generate <key>   auto generates a 12 character random value
  list             lists all available keys
  search <key>     lists all keys matching containing partial regex <key>

Example:
  - save key github.com
      kb save github.com

  - get key github.com
      kb get github.com
```
