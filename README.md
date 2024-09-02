# Purpose

A file encrypter / decrypter project, written in Go.

# Usage

## Build

Run the following from the root of the directory

```
go build -o app src/*.go
```

## Encrypt

To encryt, first build the project (shown above) then run the following:

```
./app encrypt /path/to/file
```

## Decrypt

To decrypt, run the following:

```
./app decrypt /path/to/file
```
