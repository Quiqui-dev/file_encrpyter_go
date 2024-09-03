# Purpose

A file encrypter / decrypter project, written in Go.

# Usage

## Build

Run the following from the root of the directory

```
go build -o app src/*.go
```

## Encrypt File

To encryt a file, first build the project (shown above) then run the following:

```
./app -e -f "filename" -p "/path/to/file" -t "encryption_string"
```

## Encrypt all files in a directory

To encryt all files in a directory, first build the project (shown above) then run the following:

```
./app -e -p "/path/to/files" -t "encryption_string"
```

## Decrypt a file

To decrypt a file, run the following:

```
./app -d -f "filename" -p "/path/to/file" -t "encryption_string"
```

## Decrypt all files in a directory

To decrypt all files in a directory, run the following:

```
./app -d -p "/path/to/files" -t "encryption_string"
```
