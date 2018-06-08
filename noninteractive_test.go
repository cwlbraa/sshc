package sshc_test

import (
	"github.com/cwlbraa/sshc/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("noninteractively,", func() {
	var (
		testServer *fakes.FakeSSHD
	)

	BeforeEach(func() {
		testServer = fakes.NewFakeSSHD()
	})

	JustBeforeEach(func() {
		err := testServer.StartListen()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		testServer.StopListen()
	})

	It("can SSH without authentication", func() {
		// run some command that doesn't need a tty
	})

	It("can SSH with key authentication", func() {
	})
})
