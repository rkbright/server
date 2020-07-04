// Package ssh connection
package main

import (
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	user := os.Getenv("USER")
	home := os.Getenv("HOME")
	host := os.Getenv("HOST")

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			publicKey(home + "/.ssh/id_rsa"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		panic(err)
	}

	defer conn.Close()
}

func publicKey(path string) ssh.AuthMethod {
	key, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}
	return ssh.PublicKeys(signer)
}
