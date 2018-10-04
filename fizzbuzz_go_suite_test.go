package fizzbuzz_go_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFizzbuzzGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FizzbuzzGo Suite")
}
