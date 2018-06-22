package sshc

import (
	"bytes"
	"io"
)

type Host struct {
	Name string
	Port int
	// TODO: jumpHost
}

// will execute with sh -c "shellCommand"
func (h Host) Command(shellCommand string) (io.Reader, io.Reader, error) {
	var stdout, stderr bytes.Buffer
	return &stdout, &stderr, h.CollectedCommand(shellCommand, &stdout, &stderr)
}

func (h Host) CollectedCommand(shellCommand string, stdout io.Writer, stderr io.Writer) error {
	return nil
}
