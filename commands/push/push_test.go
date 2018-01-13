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
)

var _ = Describe("PushPlan", func() {
	var (
		mockCtrl       *gomock.Controller
		manifestReader *MockManifestReader
		push           Push
		options        helpers.Options
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		manifestReader = NewMockManifestReader(mockCtrl)
		options = helpers.Options{
			ManifestPath: "/path/to/manifest.yml",
			Postfix:      "yolo",
			Domain:       "apps.cf.com",
		}
		push = NewPush(manifestReader, options)
	})

	It("Returns an error when manifest at path does not exist", func() {
		manifestReader.EXPECT().Read(options.ManifestPath).
			Return(manifest.Application{}, errors.New("Yolo"))

		app, err := push.PushPlan()

		Expect(err).To(HaveOccurred())
		Expect(app).To(Equal(Plan{}))

	})

	It("Creates a plan for a push", func() {
		application := manifest.Application{
			Name: "my-test-app",
		}

		appName := application.Name + "-" + options.Postfix

		manifestReader.EXPECT().Read(options.ManifestPath).
			Return(application, nil)

		plan, err := push.PushPlan()

		expected := Plan{
			Cmds: []Cmd{
				CfCmd{[]string{"push", appName, "-f", options.ManifestPath, "-n", appName, "-d", options.Domain}},
			},
		}

		Expect(err).To(Not(HaveOccurred()))
		Expect(plan).To(Equal(expected))
	})
})
