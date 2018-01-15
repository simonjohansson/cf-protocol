package delete_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/simonjohansson/cf-protocol/mocks"
	. "github.com/simonjohansson/cf-protocol/command"
	. "github.com/simonjohansson/cf-protocol/commands/delete"
	"code.cloudfoundry.org/cli/util/manifest"
	"github.com/golang/mock/gomock"
	"code.cloudfoundry.org/cli/cf/errors"
	"github.com/simonjohansson/cf-protocol/helpers"
)

var _ = Describe("Delete Plan", func() {
	var (
		mockCtrl       *gomock.Controller
		manifestReader *MockManifestReader
		delete         Delete
		options        helpers.Options
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		manifestReader = NewMockManifestReader(mockCtrl)
		options = helpers.Options{
			ManifestPath: "/path/to/manifest.yml",
			Postfix:      "yolo",
		}
		delete = NewDelete(manifestReader, options)
	})

	It("Returns an error when manifest at path does not exist", func() {
		manifestReader.EXPECT().Read(options.ManifestPath).
			Return(manifest.Application{}, errors.New("Yolo"))

		app, err := delete.DeletePlan()

		Expect(err).To(HaveOccurred())
		Expect(app).To(Equal(Plan{}))

	})

	It("Creates a plan for a delete", func() {
		application := manifest.Application{
			Name: "my-test-app",
		}
		manifestReader.EXPECT().Read(options.ManifestPath).
			Return(application, nil)
		appName := application.Name + "-" + options.Postfix

		plan, err := delete.DeletePlan()

		expected := Plan{
			Cmds: []Cmd{
				CfCmd{[]string{"delete", appName, "-f"}},
			},
		}

		Expect(err).To(Not(HaveOccurred()))
		Expect(plan).To(Equal(expected))
	})
})
