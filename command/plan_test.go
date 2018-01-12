package command

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Command")
}

var _ = Describe("Plan", func() {
	It("produces a plan as strings to be printed out", func() {
		plan := Plan{
			[]Cmd{
				Cmd{[]string{"a", "b", "c"}},
				Cmd{[]string{"1", "2", "3"}},
			},
		}

		Expect(plan.Printable()).To(Equal(
			[]string{
				"cf a b c",
				"cf 1 2 3",
			}))
	})
})
