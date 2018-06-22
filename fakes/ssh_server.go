package fakes

import (
	"github.com/gliderlabs/ssh"
)

type FakeSSHD struct {
	RealSSHD *ssh.Server
	errChan  chan error
}

func NewFakeSSHD() *FakeSSHD {
	return &FakeSSHD{
		RealSSHD: &ssh.Server{},
		errChan:  make(chan error),
	}
}

func (f *FakeSSHD) StartListen() error {
	// TODO seperate listen and serve for fast-fail errs
	go func() {
		f.errChan <- f.RealSSHD.ListenAndServe()
	}()

	return nil
}

func (f *FakeSSHD) StopListen() error {
	err := f.RealSSHD.Close()
	if err != nil {
		return err
	}

	return <-f.errChan
}

// TODO random available port
