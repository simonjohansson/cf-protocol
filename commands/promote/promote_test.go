package promote_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/simonjohansson/cf-protocol/mocks"
	. "github.com/simonjohansson/cf-protocol/commands/promote"
	"github.com/golang/mock/gomock"
	"github.com/simonjohansson/cf-protocol/helpers"
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
		options = helpers.Options{}
		promote = NewPromote(cliConnection, options)

	})

	It("Returns an error when manifest at path does not exist", func() {

	})
})
