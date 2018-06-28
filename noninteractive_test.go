package sshc_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/cwlbraa/sshc"
	"github.com/cwlbraa/sshc/mocks"
	sshs "github.com/gliderlabs/ssh"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"golang.org/x/crypto/ssh"
)

var _ = FDescribe("noninteractively,", func() {
	var (
		testServer             *mocks.MockSSHD
		host                   sshc.Host
		confirmRemoteExecution func()
	)

	BeforeEach(func() {
		testServer = mocks.NewMockSSHD()
	})

	JustBeforeEach(func() {
		err := testServer.Start()
		Expect(err).NotTo(HaveOccurred())

		host = sshc.Host{
			Name: "localhost",
			Port: testServer.Port(),
		}

		confirmRemoteExecution = func() {
			output := gbytes.NewBuffer()
			err := host.Command("rev <<< 123", output, output)
			Expect(output).To(gbytes.Say("321"))
			Expect(err).NotTo(HaveOccurred())
		}
	})

	AfterEach(func() {
		err := testServer.Stop()
		Expect(err).NotTo(HaveOccurred())
	})

	It("can execute shell commands remotely", func() {
		confirmRemoteExecution()
	})

	Context("with key authentication", func() {
		var key *rsa.PrivateKey
		BeforeEach(func() {
			var err error
			key, err = rsa.GenerateKey(rand.Reader, 2048)
			Expect(err).NotTo(HaveOccurred())
			clientKey, err := ssh.NewPublicKey(key.Public())
			Expect(err).NotTo(HaveOccurred())

			testServer.RealSSHD.PublicKeyHandler = func(ctx sshs.Context, receivedKey sshs.PublicKey) bool {
				return Expect(receivedKey).To(Equal(clientKey))
			}
		})

		It("can execute shell commands remotely", func() {
			err := host.ParsePrivateKey(pem.EncodeToMemory(&pem.Block{
				Type:  "RSA PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(key),
			}))
			Expect(err).NotTo(HaveOccurred())
			confirmRemoteExecution()
		})
	})
})
