package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/term"
)


func validateFile(file string) bool {

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
}


func getPassword() []byte {
	fmt.Print("enter password: ")

	password, _ := term.ReadPassword(0)

	return password
}

func validatePassword(pass1 []byte, pass2 []byte) bool {

	return bytes.Equal(pass1, pass2)
}