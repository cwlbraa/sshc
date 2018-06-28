package sshc_test

import (
	"github.com/cwlbraa/sshc/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("interactively,", func() {
	var (
		testServer *mocks.MockSSHD
	)

	BeforeEach(func() {
		testServer = mocks.NewMockSSHD()
	})

	JustBeforeEach(func() {
		err := testServer.Start()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		testServer.Stop()
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
