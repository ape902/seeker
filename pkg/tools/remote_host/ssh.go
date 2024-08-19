package remote_host

import (
	"time"

	"golang.org/x/crypto/ssh"
)

const (
	PASSWORD AuthModel = iota + 1
	PUBLICKEY
)

type (
	AuthModel int8

	SSHClientConfig struct {
		AuthModel      AuthModel
		HostAddr       string
		Username       string
		Authentication string
		Timeout        time.Duration
	}
)

func NewSSHDial(addr, username, auth string, authmode int8) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		Timeout:         time.Second * time.Duration(5),
		User:            username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	switch authmode {
	case int8(PASSWORD):
		config.Auth = []ssh.AuthMethod{ssh.Password(auth)}
	case int8(PUBLICKEY):
		publicKey, err := ssh.ParsePrivateKey([]byte(auth))
		if err != nil {
			return nil, err
		}

		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(publicKey)}
	}

	return ssh.Dial("tcp", addr, config)
}
