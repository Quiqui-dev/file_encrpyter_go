package main

import (
	"fmt"
	"os"

	"github.com/Quiqui-dev/file_encrpyter_go/src/decrypt"
	"github.com/Quiqui-dev/file_encrpyter_go/src/encrypt"
)

func main() {

	app := createApp()

	app.Run(os.Args)

}

func handleEncryptSingle(path_to_act string, password []byte) {

	if !validateFile(path_to_act) {
		fmt.Println(path_to_act)
	}

	fmt.Printf("\nEncryption started on %s", path_to_act)

	encrypt.Encrpyt(path_to_act, password)
	fmt.Println("\nfile successfully protected")
}

func handleEncryptDirectory(path_to_act string, password []byte) {

	// get all files in a dir

	files, err := os.ReadDir(path_to_act)

	if err != nil {
		panic(err.Error())
	}

	
	for _, file := range files {
		path_to_act_copy := path_to_act

		path_to_act_copy = path_to_act_copy + file.Name()

		handleEncryptSingle(path_to_act_copy, password)
	}
}

func handleDecryptSingle(path_to_act string, password []byte) {

	if !validateFile(path_to_act) {
		panic("File not found")
	}


	fmt.Printf("\nDecryption started on %s", path_to_act)

	decrypt.Decrpyt(path_to_act, password)
	fmt.Println("\nfile successfully decrypted")
}

func handleDecryptDirectory(path_to_act string, password []byte) {

	// get all files in a dir

	files, err := os.ReadDir(path_to_act)

	if err != nil {
		panic(err.Error())
	}

	
	for _, file := range files {
		path_to_act_copy := path_to_act

		path_to_act_copy = path_to_act_copy + file.Name()

		handleDecryptSingle(path_to_act_copy, password)
	}
}