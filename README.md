
**kb** is small a command line password manager.

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

* set the key github.com (this will prompt you for a pasword)
  ```text
  kb set github.com
  ```
* get the key github.com
  ```text
  kb get github.com
  ```
* list all keys
  ```text
  kb ls
  ```
* search for the key github
   ```text
   kb search github
   ```

View the ciphertext in `~/.kb/<keyname>` 
```text
$ cat ~/.kb/test
BEGIN KEYBASE SALTPACK ENCRYPTED MESSAGE. kiOUtMhcc4NXXRb XMxIdgQyljvmpjs wcr3T7EQMTg
C6o7 FW06TvfusBBSg2Y ZD74tOWYqzvqN4c E92x5CJwVQAsHNn qMGqb05pPMuqgfE hHEE3RadJ7scDrl
EEAdP9mipF0wSeX MwdyLjBYHUJdGYs 0Jh9A8k2xEtBQcD uroiYpkCOvkFckT Vimw6Lxh613J7Hk XFXCl
sAf15kd0RI EhVS456ulYeQW8r tylMbrvixB0cVbX b29wBDd70s2TYY6 Bwd7qkrTpHZQ9qZ svh559yYsg
ebMqi VHUkwcXwdhWBvGz 2lN85H9xYOADpji fdagS4kdF97TRqA n6M6FAEdXCao4qu onrBtvmehwNo4Wl
 ywsF7KrDSrnuwBk bcTJNYGeJ2y2xnM 6Ha3DoEwcFWhbS0 fjmdtZ1C1cejMIR tElYbVWNUAZq1rj hE3K
Nk28gomZVcz hBw7acF6rmDWmJU 3qqfU80FCh9tx4U ZHPu2ZJjtF5r50x JvIAZrOB5lXhjKP abztGFnfo
ROFyhO 2Yk3rF5tbb54pD5 y8VtVQY71WQhSP2 ijaue4hqHMsAqfk 5g0UcMSkMW1xHfE oKpdkThJvDYKEu
m xM1tGEJ8DoMwXYr ECN9vmcQsGmU2tN Ge9hE04mzk0YPgP X1Mc0mUHZh5zjUG iJX4eIq9PAC4AUo yt2
OcTiljHD4G7g LQQqWR0xy9MNHn7 5q5NO7RFkUFIiJX eQ2A7MxRPbKDojT BHLxIDTvAGrP16d 29hJEQvA
H7b3gvT qFOrlmsKn323jSd RDkGx9Z903nApeu VpVxpYtigxUvsdl ylHUe. END KEYBASE SALTPACK E
NCRYPTED MESSAGE.
```

