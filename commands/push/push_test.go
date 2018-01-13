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
)

var _ = Describe("PushPlan", func() {
	var (
		mockCtrl       *gomock.Controller
		manifestReader *MockManifestReader
		logger         *MockLogger
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		manifestReader = NewMockManifestReader(mockCtrl)
		logger = NewMockLogger(mockCtrl)

		logger.EXPECT().Info(gomock.Any()).AnyTimes()
	})

	It("Returns an error when manifest at path does not exist", func() {
		postfix := "asdf"
		domain := "apps.dc.springernature.io"
		manifestPath := "path/to/manifest.yml"

		manifestReader.EXPECT().Read(manifestPath).
			Return(manifest.Application{}, errors.New("Yolo"))

		app, err := PushPlan(manifestPath, postfix, domain, logger, manifestReader)

		Expect(err).To(HaveOccurred())
		Expect(app).To(Equal(Plan{}))

	})

	It("Creates a plan for a push", func() {
		application := manifest.Application{
			Name: "my-test-app",
		}

		postfix := "asdf"
		domain := "apps.dc.springernature.io"
		manifestPath := "path/to/manifest.yml"

		appName := application.Name + "-" + postfix

		manifestReader.EXPECT().Read(manifestPath).
			Return(application, nil)

		plan, err := PushPlan(manifestPath, postfix, domain, logger, manifestReader)

		expected := Plan{
			Cmds: []Cmd{
				CfCmd{[]string{"push", appName, "-f", manifestPath, "-n", appName, "-d", domain}},
			},
		}

		Expect(err).To(Not(HaveOccurred()))
		Expect(plan).To(Equal(expected))
	})
})
