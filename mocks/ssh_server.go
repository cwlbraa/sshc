package mocks

import (
	"net"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"github.com/gliderlabs/ssh"
)

type MockSSHD struct {
	listener net.Listener
	RealSSHD *ssh.Server
	errChan  chan error
}

func baseHandler(sess ssh.Session) {
	if len(sess.Command()) < 1 {
		// TODO: login shell
		return
	}
	cmd := exec.Command(sess.Command()[0], sess.Command()[1:]...)
	cmd.Stdin = sess
	cmd.Stdout = sess
	cmd.Stderr = sess.Stderr()

	err := cmd.Run()
	if err == nil {
		err = sess.Exit(0)
		if err != nil {
			panic(err)
		}
		return
	}

	if exiterr, ok := err.(*exec.ExitError); ok {
		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			err = sess.Exit(status.ExitStatus())
			if err != nil {
				panic(err)
			}
			return
		}
	}
	panic(err)
}

func NewMockSSHD() *MockSSHD {
	return (&MockSSHD{
		RealSSHD: &ssh.Server{},
		errChan:  make(chan error),
	}).Handle(baseHandler)
}

func (m *MockSSHD) Start() error {
	var err error
	m.listener, err = net.Listen("tcp", "localhost:")
	if err != nil {
		return err
	}

	go m.Serve()

	return nil
}
func (m *MockSSHD) Serve() {
	err := m.RealSSHD.Serve(m.listener)
	if !strings.Contains(err.Error(), "use of closed network connection") {
		m.errChan <- err
	}
	close(m.errChan)
}

func (m *MockSSHD) Addr() string {
	if m.listener == nil {
		panic("can't get addr until you've started listening")
	}

	return m.listener.Addr().String()
}

func (m *MockSSHD) Port() int {
	res, err := strconv.Atoi(strings.Split(m.Addr(), ":")[1])
	if err != nil {
		panic(err)
	}
	return res
}

func (m *MockSSHD) Stop() error {
	err := m.listener.Close()
	if err != nil {
		return err
	}
	err = m.RealSSHD.Close()
	if err != nil {
		return err
	}

	return <-m.errChan
}

func (m *MockSSHD) Handle(handler ssh.Handler) *MockSSHD {
	m.RealSSHD.Handler = m.spyHandler(handler)
	return m
}

func (m *MockSSHD) spyHandler(handler ssh.Handler) ssh.Handler {
	return handler
}
