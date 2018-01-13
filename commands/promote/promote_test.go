package promote_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/simonjohansson/cf-protocol/mocks"
	"github.com/golang/mock/gomock"
)

var _ = Describe("PushPlan", func() {
	var (
		mockCtrl       *gomock.Controller
		manifestReader *MockManifestReader
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		manifestReader = NewMockManifestReader(mockCtrl)
	})

	It("Returns an error when manifest at path does not exist", func() {

	})
})
