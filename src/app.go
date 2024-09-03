package main

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
)

func createApp() *cli.App {
	app := &cli.App{
		Name: "Encrypt / Decrpyt Tool",
		Usage: "Small tool to encrypt or decrypt a collection of one or more files",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "file",
				Aliases: []string{"f"},
				Usage: "The name of the file to act on\n if left blank, it will act on the entire dir",
				Required: false,
			},
			&cli.StringFlag{
				Name: "directory",
				Aliases: []string{"p"},
				Usage: "The name of the directory to act on",
				Required: false,
			},
			&cli.StringFlag{
				Name: "password",
				Aliases: []string{"t"},
				Usage: "The token to use to encrypt or decrypt",
				Required: true,
			},
			&cli.BoolFlag{
				Name: "encrypt",
				Aliases: []string{"e"},
				Usage: "The flag which enables encryption",
				Required: false,
			},
			&cli.BoolFlag{
				Name: "decrypt",
				Aliases: []string{"d"},
				Usage: "The flag which enables decryption",
				Required: false,
			},
		},
		Action: appAction,
	}

	return app
}


func appAction(c *cli.Context) error {

	directory_flag := c.String("directory")
	if directory_flag == "" {
		fmt.Println("no directory option selected")
	}

	password_flag := c.String("password")

	if password_flag == "" {
		fmt.Println("no password option selected")
	}

	file_flag := c.String("file")
	if file_flag == "" {
		log.Println("No file provided, acting on directory")
	}

	path_to_act := directory_flag + "/" + file_flag


	decrpyt_flag := c.Bool("decrypt") 
	encrpyt_flag := c.Bool("encrypt")

	if !decrpyt_flag && !encrpyt_flag {
		fmt.Println("no decrypt or encrypt option selected")
	}

	password := []byte(password_flag)
	
	if encrpyt_flag && file_flag != "" {
		handleEncryptSingle(path_to_act, password)
	} else if encrpyt_flag && file_flag == "" {
		handleEncryptDirectory(path_to_act, password)
	} else if decrpyt_flag && file_flag != "" {
		handleDecryptSingle(path_to_act, password)
	} else if decrpyt_flag && file_flag == "" {
		handleDecryptDirectory(path_to_act, password)
	}

	return nil
}