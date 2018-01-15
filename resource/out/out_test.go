package out_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/simonjohansson/cf-protocol/resource/out"
	. "github.com/simonjohansson/cf-protocol/command"
)

var _ = Describe("Out", func() {

	var (
		sourceRoot   string
		input        Input
		concourseEnv ConcourseEnv
		out          Out
	)

	BeforeEach(func() {
		sourceRoot = "/source/root"
		input = Input{}
		concourseEnv = ConcourseEnv{}
		out = NewOut(sourceRoot, &input, &concourseEnv)
	})

	Context("When missing source and params", func() {
		It("errors out", func() {
			plan, err := out.OutPlan()
			Expect(err).To(HaveOccurred())
			Expect(plan).To(Equal(Plan{}))
		})
	})

	Context("When required data is passed to the plan", func() {
		It("creates a plan where the first CMD is a cd to the sourceRoot/params.dir", func() {
			input = Input{
				Params: Params{
					Dir: "/some/path",
				},
			}

			plan, err := out.OutPlan()
			Expect(err).To(Not(HaveOccurred()))
			Expect(plan.Cmds[0]).To(Equal(CliCmd{
				[]string{"cd", "/source/root/some/path"},
			}))
		})

		It("creates a plan where the second CMD is a cf login", func() {
			input = Input{
				Source: Source{
					"api",
					"username",
					"password",
					"org",
					"space",
				},
				Params: Params{
					Dir: "/some/path",
				},
			}

			plan, err := out.OutPlan()
			Expect(err).To(Not(HaveOccurred()))
			Expect(plan.Cmds[1]).To(Equal(CliCmd{
				[]string{"cf", "login",
					"-a", input.Source.Api,
					"-u", input.Source.Username,
					"-p", input.Source.Password,
					"-o", input.Source.Org,
					"-s", input.Source.Password},
			}))
		})

		It("creates a plan where the third CMD is a cf protocol-*", func() {
			input = Input{
				Source: Source{
					"api",
					"username",
					"password",
					"org",
					"space",
				},
				Params: Params{
					Dir:          "/some/path",
					Cmd:          "push",
					ManifestPath: "manifest.yml",
				},
			}
			concourseEnv = ConcourseEnv{
				BuildName: "1337",
			}

			plan, err := out.OutPlan()
			Expect(err).To(Not(HaveOccurred()))
			Expect(plan.Cmds[2]).To(Equal(CliCmd{
				[]string{"cf", "protocol-push", "-manifest", "manifest.yml", "-domain", "domain.io", "-postfix", "1337"},
			}))
		})

	})
})
