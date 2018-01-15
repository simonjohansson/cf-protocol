package cleanup_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/simonjohansson/cf-protocol/mocks"
	. "github.com/simonjohansson/cf-protocol/command"
	. "github.com/simonjohansson/cf-protocol/commands/cleanup"
	"github.com/golang/mock/gomock"
	"github.com/simonjohansson/cf-protocol/helpers"
	"errors"
	"code.cloudfoundry.org/cli/util/manifest"
	"code.cloudfoundry.org/cli/plugin/models"
)

var _ = Describe("Cleanup plan", func() {
	var (
		mockCtrl       *gomock.Controller
		manifestReader *MockManifestReader
		cliConnection  *MockCliConnection
		cleanup        Cleanup
		options        helpers.Options
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		manifestReader = NewMockManifestReader(mockCtrl)
		cliConnection = NewMockCliConnection(mockCtrl)
		options = helpers.Options{
			ManifestPath: "/path/to/manifest.yml",
			Postfix:      "1337",
		}
		cleanup = NewCleanup(cliConnection, manifestReader, options)

	})

	It("Returns an error when manifest at path does not exist", func() {
		e := errors.New("Yolo")
		manifestReader.EXPECT().Read(options.ManifestPath).
			Return(manifest.Application{}, e)

		app, err := cleanup.CleanupPlan()

		Expect(err).To(Equal(e))
		Expect(app).To(Equal(Plan{}))
	})

	Context("When there is no old apps with the same postfix", func() {
		It("does nothing", func() {
			application := manifest.Application{
				Name:   "my-app",
				Routes: []string{"route-a.domain1.com", "route-b.subdomain.domain2.com"},
			}
			manifestReader.EXPECT().Read(options.ManifestPath).
				Return(application, nil)

			appName := application.Name + "-" + options.Postfix
			cliConnection.EXPECT().GetApps().Return([]plugin_models.GetAppsModel{
				plugin_models.GetAppsModel{Name: "other-app-1"},
				plugin_models.GetAppsModel{Name: appName, State: "started"},
				plugin_models.GetAppsModel{Name: "yolo"},
			}, nil)

			plan, err := cleanup.CleanupPlan()

			Expect(err).To(Not(HaveOccurred()))
			Expect(plan).To(Equal(Plan{[]Cmd{}}))
		})
	})

	Context("When there is one old app with the same name", func() {
		It("stops the old running app", func() {
			application := manifest.Application{
				Name:   "my-app",
				Routes: []string{"route-a.domain1.com", "route-b.subdomain.domain2.com"},
			}
			manifestReader.EXPECT().Read(options.ManifestPath).
				Return(application, nil)

			appName := application.Name + "-" + options.Postfix
			oldAppName := application.Name + "-" + "420"
			cliConnection.EXPECT().GetApps().Return([]plugin_models.GetAppsModel{
				plugin_models.GetAppsModel{Name: "other-app-1", State: "started"},
				plugin_models.GetAppsModel{Name: appName, State: "started"},
				plugin_models.GetAppsModel{Name: oldAppName, State: "started"},
				plugin_models.GetAppsModel{Name: "yolo", State: "stopped"},
			}, nil)

			plan, err := cleanup.CleanupPlan()

			Expect(err).To(Not(HaveOccurred()))
			Expect(plan).To(Equal(Plan{[]Cmd{
				CfCmd{
					[]string{"stop", oldAppName},
				},
			}}))
		})

		It("does nothing if its stopped already", func() {
			application := manifest.Application{
				Name:   "my-app",
				Routes: []string{"route-a.domain1.com", "route-b.subdomain.domain2.com"},
			}
			manifestReader.EXPECT().Read(options.ManifestPath).
				Return(application, nil)

			appName := application.Name + "-" + options.Postfix
			oldAppName := application.Name + "-" + "420"
			cliConnection.EXPECT().GetApps().Return([]plugin_models.GetAppsModel{
				plugin_models.GetAppsModel{Name: "other-app-1", State: "started"},
				plugin_models.GetAppsModel{Name: appName, State: "started"},
				plugin_models.GetAppsModel{Name: oldAppName, State: "stopped"},
				plugin_models.GetAppsModel{Name: "yolo", State: "stopped"},
			}, nil)

			plan, err := cleanup.CleanupPlan()

			Expect(err).To(Not(HaveOccurred()))
			Expect(plan).To(Equal(Plan{[]Cmd{}}))
		})
	})

	Context("When there is multiple old app with the same name", func() {
		It("Stops the next to latest app and deletes the other ones", func() {
			application := manifest.Application{
				Name:   "my-app",
				Routes: []string{"route-a.domain1.com", "route-b.subdomain.domain2.com"},
			}
			manifestReader.EXPECT().Read(options.ManifestPath).
				Return(application, nil)

			appName := application.Name + "-" + options.Postfix
			oldAppName1 := application.Name + "-" + "100"
			oldAppName2 := application.Name + "-" + "200"
			oldAppName3 := application.Name + "-" + "300"
			cliConnection.EXPECT().GetApps().Return([]plugin_models.GetAppsModel{
				plugin_models.GetAppsModel{Name: "other-app-1", State: "started"},
				plugin_models.GetAppsModel{Name: appName, State: "started"},
				plugin_models.GetAppsModel{Name: oldAppName2, State: "started"},
				plugin_models.GetAppsModel{Name: oldAppName1, State: "stopped"},
				plugin_models.GetAppsModel{Name: oldAppName3, State: "started"},
				plugin_models.GetAppsModel{Name: "yolo", State: "stopped"},
			}, nil)

			plan, err := cleanup.CleanupPlan()

			Expect(err).To(Not(HaveOccurred()))
			Expect(plan).To(Equal(Plan{[]Cmd{
				CfCmd{[]string{"stop", oldAppName3}},
				CfCmd{[]string{"delete", oldAppName2, "-f"}},
				CfCmd{[]string{"delete", oldAppName1, "-f"}},
			}}))

		})
	})
})
