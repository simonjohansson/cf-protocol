package conformance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/golang/mock/gomock"
	. "github.com/simonjohansson/cf-protocol/mocks"
	. "github.com/simonjohansson/cf-protocol/commands/conformance"
	"io/ioutil"
	"bytes"
	"net/http"
)

func makeResponse(data string, returnCode int) *http.Response {
	return &http.Response{
		Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(data))),
		StatusCode: returnCode,
	}
}

var _ = Describe("Conformance", func() {
	var (
		mockCtrl   *gomock.Controller
		httpClient *MockHttpClient
		logger     *MockLogger
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		httpClient = NewMockHttpClient(mockCtrl)
		logger = NewMockLogger(mockCtrl)

		logger.EXPECT().Info(gomock.Any()).AnyTimes()
	})

	It("Works when both internal/version is correct and internal/status is 200", func() {
		appUrl := "https://app-under-test.cf.io"

		httpClient.EXPECT().Get(appUrl + "/internal/version").
			Return(makeResponse(`{"revision": "a840f1ae2122c3540a47541887eccff04aaeb212"}`, 200), nil, )

		httpClient.EXPECT().Get(appUrl + "/internal/status").
			Return(makeResponse(``, 200), nil, )

		err := Conformance(appUrl, httpClient, logger)

		Expect(err).To(BeNil())

	})
	//
	It("Errors when non 200 for version", func() {
		appUrl := "https://app-under-test.cf.io"

		httpClient.EXPECT().Get(appUrl + "/internal/version").
			Return(makeResponse(``, 500), nil, )

		httpClient.EXPECT().Get(appUrl + "/internal/status").
			Return(makeResponse(``, 200), nil, )

		err := Conformance(appUrl, httpClient, logger)

		Expect(err).To(Not(BeNil()))
	})

	It("Errors when non valid json returned for version", func() {
		appUrl := "https://app-under-test.cf.io"

		httpClient.EXPECT().Get(appUrl + "/internal/version").
			Return(makeResponse(`{"im borked": ...'`, 500), nil, )

		httpClient.EXPECT().Get(appUrl + "/internal/status").
			Return(makeResponse(``, 200), nil, )

		err := Conformance(appUrl, httpClient, logger)

		Expect(err).To(Not(BeNil()))
	})

	It("Errors when non 200 for /internal/status", func() {
		appUrl := "https://app-under-test.cf.io"

		httpClient.EXPECT().Get(appUrl + "/internal/version").
			Return(makeResponse(`{"revision": "a840f1ae2122c3540a47541887eccff04aaeb212"}`, 200), nil, )

		httpClient.EXPECT().Get(appUrl + "/internal/status").
			Return(makeResponse(``, 500), nil, )

		err := Conformance(appUrl, httpClient, logger)

		Expect(err).To(Not(BeNil()))
	})
})
