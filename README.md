## KBPass

Command line password manager using the keybase filesystem for
password storage.

#### Installation

```
go get github.com/kingishb/kbpass
```

#### Usage
First, install [keybase](https://keybase.io/).

Then, run
```
kbpass init
```

to build a keystore in your personal private keybase folder.

Save passwords like

```
kbpass save github
```

Retrieve passwords with

```
kbpass get github
```

Share passwords with

```
keybase share github jdsallinger
```


