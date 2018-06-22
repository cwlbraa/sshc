package sshc_test

import (
	"github.com/cwlbraa/sshc"
	"github.com/cwlbraa/sshc/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("noninteractively,", func() {
	var (
		testServer *fakes.FakeSSHD
		host       sshc.Host
	)

	BeforeEach(func() {
		testServer = fakes.NewFakeSSHD()
		host = sshc.Host{Name: "localhost", Port: 22}
	})

	JustBeforeEach(func() {
		err := testServer.StartListen()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		testServer.StopListen()
	})

	FIt("can SSH without authentication", func() {
		output := gbytes.NewBuffer()
		err := host.CollectedCommand("rev <<< 123", output, output)
		Expect(output).To(gbytes.Say("321"))
		Expect(err).NotTo(HaveOccurred())
	})

	It("can SSH with key authentication", func() {
	})
})
