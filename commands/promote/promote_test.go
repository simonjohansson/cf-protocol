package promote_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/simonjohansson/cf-protocol/mocks"
	. "github.com/simonjohansson/cf-protocol/commands/promote"
	. "github.com/simonjohansson/cf-protocol/command"
	"github.com/golang/mock/gomock"
	"github.com/simonjohansson/cf-protocol/helpers"
	"errors"
	"code.cloudfoundry.org/cli/util/manifest"
	"code.cloudfoundry.org/cli/plugin/models"
)

var _ = Describe("Promote plan", func() {
	var (
		mockCtrl       *gomock.Controller
		manifestReader *MockManifestReader
		cliConnection  *MockCliConnection
		promote        Promote
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
		promote = NewPromote(cliConnection, manifestReader, options)

	})

	It("Returns an error when manifest at path does not exist", func() {
		e := errors.New("Yolo")
		manifestReader.EXPECT().Read(options.ManifestPath).
			Return(manifest.Application{}, e)

		app, err := promote.PromotePlan()

		Expect(err).To(Equal(e))
		Expect(app).To(Equal(Plan{}))

	})

	Context("When there is no old apps", func() {
		It("binds the routes", func() {
			application := manifest.Application{
				Name:   "my-app",
				Routes: []string{"route-a.domain1.com", "route-b.subdomain.domain2.com"},
			}
			manifestReader.EXPECT().Read(options.ManifestPath).
				Return(application, nil)

			cliConnection.EXPECT().GetApps().Return([]plugin_models.GetAppsModel{}, nil)

			appName := application.Name + "-" + options.Postfix
			app, err := promote.PromotePlan()

			Expect(err).To(Not(HaveOccurred()))
			Expect(app).To(Equal(
				Plan{
					[]Cmd{
						CfCmd{[]string{"map-route", appName, "domain1.com", "--hostname", "route-a"}},
						CfCmd{[]string{"map-route", appName, "subdomain.domain2.com", "--hostname", "route-b"}},
					}}))
		})
	})

	Context("When there is one old app with the same postfix", func() {
		It("binds the routes and stops the old running app", func() {
			application := manifest.Application{
				Name:   "my-app",
				Routes: []string{"route-a.domain1.com", "route-b.subdomain.domain2.com"},
			}
			manifestReader.EXPECT().Read(options.ManifestPath).
				Return(application, nil)

			oldAppName := application.Name + "-" + "420"
			cliConnection.EXPECT().GetApps().Return(
				[]plugin_models.GetAppsModel{
					plugin_models.GetAppsModel{Name: "other-app-1"},
					plugin_models.GetAppsModel{Name: oldAppName, State: "started"},
					plugin_models.GetAppsModel{Name: "random"},
				}, nil)
			appName := application.Name + "-" + options.Postfix
			app, err := promote.PromotePlan()

			Expect(err).To(Not(HaveOccurred()))
			Expect(app).To(Equal(
				Plan{
					[]Cmd{
						CfCmd{[]string{"map-route", appName, "domain1.com", "--hostname", "route-a"}},
						CfCmd{[]string{"map-route", appName, "subdomain.domain2.com", "--hostname", "route-b"}},
						CfCmd{[]string{"stop", oldAppName}},
					}}))
		})

		It("binds the routes and doesn't stop the old running app", func() {
			application := manifest.Application{
				Name:   "my-app",
				Routes: []string{"route-a.domain1.com", "route-b.subdomain.domain2.com"},
			}
			manifestReader.EXPECT().Read(options.ManifestPath).
				Return(application, nil)

			oldAppName := application.Name + "-" + "420"
			cliConnection.EXPECT().GetApps().Return(
				[]plugin_models.GetAppsModel{
					plugin_models.GetAppsModel{Name: "other-app-1"},
					plugin_models.GetAppsModel{Name: oldAppName, State: "stopped"},
					plugin_models.GetAppsModel{Name: "random"},
				}, nil)
			appName := application.Name + "-" + options.Postfix
			app, err := promote.PromotePlan()

			Expect(err).To(Not(HaveOccurred()))
			Expect(app).To(Equal(
				Plan{
					[]Cmd{
						CfCmd{[]string{"map-route", appName, "domain1.com", "--hostname", "route-a"}},
						CfCmd{[]string{"map-route", appName, "subdomain.domain2.com", "--hostname", "route-b"}},
					}}))
		})
	})

	Context("When there is one newer app with a greater postfix", func() {
		It("errors out", func() {
			application := manifest.Application{
				Name:   "my-app",
				Routes: []string{"route-a.domain1.com", "route-b.subdomain.domain2.com"},
			}
			manifestReader.EXPECT().Read(options.ManifestPath).
				Return(application, nil)

			newerAppName := application.Name + "-" + "1338"
			cliConnection.EXPECT().GetApps().Return(
				[]plugin_models.GetAppsModel{
					plugin_models.GetAppsModel{Name: "other-app-1"},
					plugin_models.GetAppsModel{Name: newerAppName, State: "started"},
					plugin_models.GetAppsModel{Name: "random"},
				}, nil)

			app, err := promote.PromotePlan()

			Expect(err).To(HaveOccurred())
			Expect(app).To(Equal(Plan{}))
		})
	})

	Context("When there is one app with the same postfix", func() {
		It("it maps the routes but doesn't stop the app in the end", func() {
			application := manifest.Application{
				Name:   "my-app",
				Routes: []string{"route-a.domain1.com", "route-b.subdomain.domain2.com"},
			}
			manifestReader.EXPECT().Read(options.ManifestPath).
				Return(application, nil)

			appName := application.Name + "-" + options.Postfix
			cliConnection.EXPECT().GetApps().Return(
				[]plugin_models.GetAppsModel{
					plugin_models.GetAppsModel{Name: "other-app-1"},
					plugin_models.GetAppsModel{Name: appName, State: "started"},
					plugin_models.GetAppsModel{Name: "random"},
				}, nil)

			app, err := promote.PromotePlan()

			Expect(err).To(Not(HaveOccurred()))
			Expect(app).To(Equal(
				Plan{
					[]Cmd{
						CfCmd{[]string{"map-route", appName, "domain1.com", "--hostname", "route-a"}},
						CfCmd{[]string{"map-route", appName, "subdomain.domain2.com", "--hostname", "route-b"}},
					}}))
		})
	})
})
