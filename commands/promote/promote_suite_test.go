package promote_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPromote(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Promote Suite")
}
