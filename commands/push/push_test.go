package push

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	. "github.com/simonjohansson/cf-protocol/helpers"
	. "github.com/simonjohansson/cf-protocol/command"
	"code.cloudfoundry.org/cli/util/manifest"
	"errors"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Push")
}

var _ = Describe("PushPlan", func() {
	It("Returns an error when manifest at path does not exist", func() {
		manifestReader := NewMockManifestReader()
		manifestReader.SetApplication(manifest.Application{Name: "Simon"}, errors.New("Bla bla path not found"))

		postfix := "asdf"
		domain := "apps.dc.springernature.io"
		manifestPath := "path/to/manifest.yml"

		app, err := PushPlan(manifestPath, postfix, domain, NewMockLogger(), manifestReader)

		Expect(err).To(Not(BeNil()))
		Expect(app).To(Equal(Plan{}))

	})

	It("Creates a plan for a push", func() {
		application := manifest.Application{
			Name: "my-test-app",
		}

		manifestReader := NewMockManifestReader()
		manifestReader.SetApplication(application, nil)

		postfix := "asdf"
		domain := "apps.dc.springernature.io"
		manifestPath := "path/to/manifest.yml"

		appName := application.Name + "-" + postfix

		plan, err := PushPlan(manifestPath, postfix, domain, NewMockLogger(), manifestReader)

		expected := Plan{
			Cmds: []Cmd{
				Cmd{[]string{"push", appName, "-f", manifestPath, "-n", appName, "-d", domain}},
			},
		}

		Expect(err).To(BeNil())
		Expect(plan).To(Equal(expected))
	})
})
