package sshc

import (
	"bytes"
	"fmt"
	"io"

	"golang.org/x/crypto/ssh"
)

type Host struct {
	Name string
	Port int

	privateKey ssh.Signer
	// TODO: jumpHost
}

// TODO: LoadPrivateKey() < takes a path
func (h *Host) ParsePrivateKey(privateKey []byte) error {
	var err error
	h.privateKey, err = ssh.ParsePrivateKey(privateKey)
	return err
}

// will execute with sh -c "shellCommand"
func (h *Host) BufferedCommand(shellCommand string) (io.Reader, io.Reader, error) {
	var outBuffer, errBuffer bytes.Buffer
	return &outBuffer, &errBuffer, h.Command(shellCommand, &outBuffer, &errBuffer)
}

func (h *Host) Command(shellCommand string, outBuffer io.Writer, errBuffer io.Writer) error {
	config := ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	if h.privateKey != nil {
		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(h.privateKey)}
	}
	client, err := ssh.Dial("tcp", h.Addr(), &config)
	if err != nil {
		panic(err)
	}

	session, err := client.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.Stdout = outBuffer
	session.Stdout = errBuffer
	err = session.Start(fmt.Sprintf("sh -c \"%s\"", shellCommand))
	if err != nil {
		panic(err)
	}

	return session.Wait()
}

func (h *Host) Addr() string {
	return fmt.Sprintf("%s:%d", h.Name, h.Port)
}
