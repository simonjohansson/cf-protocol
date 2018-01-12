package push

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	. "github.com/simonjohansson/cf-protocol/mocks"
	. "github.com/simonjohansson/cf-protocol/command"
	"code.cloudfoundry.org/cli/util/manifest"
	"github.com/golang/mock/gomock"
	"code.cloudfoundry.org/cli/cf/errors"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Push")
}

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

		Expect(err).To(Not(BeNil()))
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

		Expect(err).To(BeNil())
		Expect(plan).To(Equal(expected))
	})
})
