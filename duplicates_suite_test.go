package duplicates

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDuplicates(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Duplicates Suite")
}
