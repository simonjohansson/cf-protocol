package delete

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
	RunSpecs(t, "Delete")
}

var _ = Describe("DeletePlan", func() {
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
		manifestPath := "path/to/manifest.yml"

		manifestReader.EXPECT().Read(manifestPath).
			Return(manifest.Application{}, errors.New("Yolo"))

		app, err := DeletePlan(manifestPath, postfix, logger, manifestReader)

		Expect(err).To(Not(BeNil()))
		Expect(app).To(Equal(Plan{}))

	})

	It("Creates a plan for a delete", func() {
		application := manifest.Application{
			Name: "my-test-app",
		}

		postfix := "asdf"
		manifestPath := "path/to/manifest.yml"

		appName := application.Name + "-" + postfix

		manifestReader.EXPECT().Read(manifestPath).
			Return(application, nil)


		plan, err := DeletePlan(manifestPath, postfix, logger, manifestReader)

		expected := Plan{
			Cmds: []Cmd{
				CfCmd{[]string{"delete", appName, "-r", "-f"}},
			},
		}

		Expect(err).To(BeNil())
		Expect(plan).To(Equal(expected))
	})
})
