package main

import (
	"fmt"
	"os"

	"github.com/Quiqui-dev/file_encrpyter_go/src/decrypt"
	"github.com/Quiqui-dev/file_encrpyter_go/src/encrypt"
	"golang.org/x/term"
)

func main() {

	if len(os.Args) < 2 { 
		handleHelp()
		os.Exit(0)
	}

	function := os.Args[1]

	switch function {
	case "help":
		handleHelp()
	case "encrypt":
		handleEncrypt()
	case "decrypt":
		handleDecrypt()
	default:
		fmt.Println("Run encrypt to encrypt a file\n Run decrypt to decrypt a file")
		os.Exit(1)
	}
}

func handleHelp() {
	fmt.Println("File encrpytion tool\n\nUsage:\n\t./exe /path/to/file\nCommands:\n\t encrpyt\t encrpyts a file given a password\n\t decrpyt \t attempts to decrypt given a password and file\n\t help \tDisplays help text")
}

func handleEncrypt() {
	if len(os.Args) < 3 {
		fmt.Println("Missing path to file")
		os.Exit(1)
	}

	file := os.Args[2]

	if !validateFile(file) {
		panic("File not found")
	}

	password := getPassword()

	fmt.Println("\nConfirm password: ")

	confrimPass, _ := term.ReadPassword(0)

	if !validatePassword(password, confrimPass) {
		fmt.Println("\nPasswords do not match. Please try again")

		password = getPassword()
	}

	fmt.Println("\nEncryption started on file")

	encrypt.Encrpyt(file, password)
	fmt.Println("\nfile successfully protected")
}

func handleDecrypt() {
	if len(os.Args) < 3 {
		fmt.Println("Missing path to file")
		os.Exit(1)
	}

	file := os.Args[2]

	if !validateFile(file) {
		panic("File not found")
	}

	password := getPassword()

	fmt.Println("\nDecryption started on file")

	decrypt.Decrpyt(file, password)
	fmt.Println("\nfile successfully decrypted")
}