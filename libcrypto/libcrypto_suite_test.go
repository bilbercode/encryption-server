package libcrypto_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLibcrypto(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Libcrypto Suite")
}
