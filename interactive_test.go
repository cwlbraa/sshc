package sshc_test

import (
	"github.com/cwlbraa/sshc/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("interactively,", func() {
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
		// run some command that needs a tty
	})

	Context("when no command is provided", func() {
		BeforeEach(func() {
		})

		It("can run a login shell", func() {
		})
	})

	Context("with key authentication", func() {
		BeforeEach(func() {
		})

		It("can allocate ttys for the executed command", func() {
		})
	})
})
