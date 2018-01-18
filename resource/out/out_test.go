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
		input = Input{

		}
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
		It("creates a plan where the first CMD is a cf login", func() {
			input = Input{
				Source: Source{
					"api",
					"username",
					"password",
				},
				Params: Params{
					Path:         "/some/path",
					Org:          "org",
					Space:        "space",
					Cmd:          "anything",
					ManifestPath: "blah",
				},
			}

			plan, err := out.OutPlan()
			Expect(err).To(Not(HaveOccurred()))
			Expect(plan.Cmds[0]).To(Equal(CliCmd{
				Args: []string{"cf", "login",
					"-a", input.Source.Api,
					"-u", input.Source.Username,
					"-p", input.Source.Password,
					"-o", input.Params.Org,
					"-s", input.Params.Space},
				Dir: "",
			}))
		})

		It("creates a plan where the second CMD is a cf protocol-push", func() {
			input = Input{
				Source: Source{
					"api",
					"username",
					"password",
				},
				Params: Params{
					Path:         "/some/path",
					Cmd:          "push",
					ManifestPath: "manifest.yml",
					Org:          "org",
					Space:        "space",
					TestDomain:   "wryyyy.io",
				},
			}
			concourseEnv = ConcourseEnv{
				BuildName: "1337",
			}

			plan, err := out.OutPlan()
			Expect(err).To(Not(HaveOccurred()))
			Expect(plan.Cmds[1]).To(Equal(CliCmd{
				Args: []string{"cf", "protocol-push",
					"-manifest", input.Params.ManifestPath,
					"-domain", input.Params.TestDomain,
					"-postfix", concourseEnv.BuildName},
				Dir: "/source/root/some/path",
			}))
		})

		It("creates a plan where the second CMD is a cf protocol-promote", func() {
			input = Input{
				Source: Source{
					"api",
					"username",
					"password",
				},
				Params: Params{
					Path:         "/some/path",
					Cmd:          "promote",
					ManifestPath: "manifest.yml",
					Org:          "org",
					Space:        "space",
				},
			}

			concourseEnv = ConcourseEnv{
				BuildName: "1337",
			}

			plan, err := out.OutPlan()
			Expect(err).To(Not(HaveOccurred()))
			Expect(plan.Cmds[1]).To(Equal(CliCmd{
				Args: []string{"cf", "protocol-promote",
					"-manifest", input.Params.ManifestPath,
					"-postfix", concourseEnv.BuildName,
				},
				Dir: "/source/root/some/path",
			}))
		})

		It("creates a plan where the second CMD is a cf protocol-cleanup", func() {
			input = Input{
				Source: Source{
					"api",
					"username",
					"password",
				},
				Params: Params{
					Path:         "/some/path",
					Cmd:          "cleanup",
					ManifestPath: "manifest.yml",
					Org:          "org",
					Space:        "space",
				},
			}

			concourseEnv = ConcourseEnv{
				BuildName: "1337",
			}

			plan, err := out.OutPlan()
			Expect(err).To(Not(HaveOccurred()))
			Expect(plan.Cmds[1]).To(Equal(CliCmd{
				Args: []string{"cf", "protocol-cleanup",
					"-manifest", input.Params.ManifestPath,
					"-postfix", concourseEnv.BuildName,
				},
				Dir: "/source/root/some/path",
			}))
		})

		It("creates a plan where the second CMD is a cf protocol-delete", func() {
			input = Input{
				Source: Source{
					"api",
					"username",
					"password",
				},
				Params: Params{
					Path:         "/some/path",
					Cmd:          "delete",
					ManifestPath: "manifest.yml",
					Org:          "org",
					Space:        "space",
				},
			}

			concourseEnv = ConcourseEnv{
				BuildName: "1337",
			}

			plan, err := out.OutPlan()
			Expect(err).To(Not(HaveOccurred()))
			Expect(plan.Cmds[1]).To(Equal(CliCmd{
				Args: []string{"cf", "protocol-delete",
					"-manifest", input.Params.ManifestPath,
					"-postfix", concourseEnv.BuildName,
				},
				Dir: "/source/root/some/path",
			}))
		})
	})
})
