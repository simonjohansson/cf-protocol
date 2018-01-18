package push_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/simonjohansson/cf-protocol/mocks"
	. "github.com/simonjohansson/cf-protocol/command"
	. "github.com/simonjohansson/cf-protocol/commands/push"
	"code.cloudfoundry.org/cli/util/manifest"
	"github.com/golang/mock/gomock"
	"code.cloudfoundry.org/cli/cf/errors"
	"github.com/simonjohansson/cf-protocol/helpers"
	"code.cloudfoundry.org/cli/plugin/models"
)

var _ = Describe("Push Plan", func() {
	var (
		mockCtrl       *gomock.Controller
		manifestReader *MockManifestReader
		cliConnection  *MockCliConnection
		push           Push
		options        helpers.Options
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		manifestReader = NewMockManifestReader(mockCtrl)
		cliConnection = NewMockCliConnection(mockCtrl)
		options = helpers.Options{
			ManifestPath: "/path/to/manifest.yml",
			Postfix:      "yolo",
			Domain:       "apps.cf.com",
		}
		push = NewPush(manifestReader, cliConnection, options)
	})

	It("Returns an error when manifest at path does not exist", func() {
		manifestReader.EXPECT().Read(options.ManifestPath).
			Return(manifest.Application{}, errors.New("Yolo"))

		app, err := push.PushPlan()

		Expect(err).To(HaveOccurred())
		Expect(app).To(Equal(Plan{}))

	})

	It("Creates a plan for a push", func() {
		spaceName := "asdfYAY"
		application := manifest.Application{
			Name: "my-test-app",
		}

		cliConnection.EXPECT().GetCurrentSpace().
			Return(plugin_models.Space{
			plugin_models.SpaceFields{
				Name: spaceName,
			},
		}, nil)

		appName := application.Name + "-" + options.Postfix
		appHostname := application.Name + "-" + spaceName + "-" + "test"

		manifestReader.EXPECT().Read(options.ManifestPath).
			Return(application, nil)

		plan, err := push.PushPlan()

		expected := Plan{
			Cmds: []Cmd{
				CfCmd{[]string{"push", appName, "-f", options.ManifestPath, "-n", appHostname, "-d", options.Domain}},
			},
		}

		Expect(err).To(Not(HaveOccurred()))
		Expect(plan).To(Equal(expected))
	})
})
