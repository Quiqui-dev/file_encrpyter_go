package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func Decrpyt(source string, password []byte) {
	
	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}

	srcFile, err := os.Open(source)

	if err != nil {
		panic(err.Error())
	}

	defer srcFile.Close()

	cipherText, err := io.ReadAll(srcFile)

	if err != nil {
		panic(err.Error())
	}

	
	key := password
	salt := cipherText[len(cipherText) - 12:]
	str := hex.EncodeToString(salt)

	nonce, err := hex.DecodeString(str)

	if err != nil {
		panic(err.Error())
	}

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)

	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	plainText, err := aesgcm.Open(nil, nonce, cipherText[:len(cipherText)-12], nil)

	if err != nil {
		panic(err.Error())
	}	

	dstFile, err := os.Create(source) 

	if err != nil {
		panic(err.Error())
	}

	defer dstFile.Close()

	_, err = dstFile.Write(plainText)

	if err != nil {
		panic(err.Error())
	}
}
