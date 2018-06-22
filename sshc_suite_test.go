package sshc_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSshc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sshc Suite")
}
