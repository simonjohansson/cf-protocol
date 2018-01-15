package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/simonjohansson/cf-protocol/command"
)

var _ = Describe("Plan", func() {
	It("produces a plan as strings to be printed out", func() {
		plan := Plan{
			[]Cmd{
				CfCmd{[]string{"a", "b", "c"}},
				CfCmd{[]string{"1", "2", "3"}},
				CliCmd{[]string{"df", "-h"}, ""},
			},
		}

		Expect(plan.Printable()).To(Equal(
			[]string{
				"cf a b c",
				"cf 1 2 3",
				"df -h",
			}))
	})
})
