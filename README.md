
#  Ash-file-Encryptor/Decryptor

  

File encryptor/decryptor with Golang

  

##  Installation

> Works perfectly on Windows 10 and Ubuntu 20.04

You can use command `go build ash.go` to complie the code to executable file

or you can download the `ash.exe` file from [Release Section](https://github.com/shimafallah/ash-file-encryptor/releases)


  

in your Command Prompt go to same directory and use `ash` command or for better experience you can add `ash.exe` to your PATH

  

##  Flags

  

`-e` Encryption Mode

`-d` Decryption Mode

  

`-f` FileName

`-p` Password

  

##  Usage

###  Encryption

`ash -e -p 123321 -f picture.png`

>it will encrypt picture.png to picture.ash with **123321** password

  

###  Decryption

`ash -d -p 123321 -f picture.ash`

>it will decrypt picture.ash to picture.png with **123321** password