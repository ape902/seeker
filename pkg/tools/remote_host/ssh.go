package remote_host

import (
	"github.com/ape902/seeker/pkg/global"
	"golang.org/x/crypto/ssh"
	"time"
)

type (
	SSHClientConfig struct {
		AuthModel      global.AuthMode
		HostAddr       string
		Username       string
		Authentication string
		Timeout        time.Duration
	}

	SSHClient struct {
		Client *ssh.Client
	}
)

func NewSSHDial(addr, username, auth string, authmode int8) (*SSHClient, error) {
	config := &ssh.ClientConfig{
		Timeout:         time.Second * time.Duration(5),
		User:            username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	switch authmode {
	case int8(global.PASSWORD):
		config.Auth = []ssh.AuthMethod{ssh.Password(auth)}
	case int8(global.PUBLICKEY):
		publicKey, err := ssh.ParsePrivateKey([]byte(auth))
		if err != nil {
			return nil, err
		}

		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(publicKey)}
	}

	dial, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, err
	}
	return &SSHClient{Client: dial}, nil
}
