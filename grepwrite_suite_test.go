package grepwrite_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGrepwrite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Grepwrite Suite")
}
