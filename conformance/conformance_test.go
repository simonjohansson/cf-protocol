package conformance

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"github.com/simonjohansson/cf-protocol/logger"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Conformance")
}

var _ = Describe("Conformance", func() {
	It("Finds revision", func() {
		appUrl := "https://app-under-test.cf.io"

		httpFetcher := NewMockHttpFetcher()
		httpFetcher.SetupTestData(appUrl+"/internal/version", `{"revision": "a840f1ae2122c3540a47541887eccff04aaeb212"}`, 200)
		httpFetcher.SetupTestData(appUrl+"/internal/status", `{}`, 200)

		err := Conformance(appUrl, httpFetcher, logger.NewMockLogger())

		Expect(err).To(BeNil())

	})

	It("Errors when non 200 for version", func() {
		appUrl := "https://app-under-test.cf.io"

		httpFetcher := NewMockHttpFetcher()
		httpFetcher.SetupTestData(appUrl+"/internal/version", ``, 500)
		httpFetcher.SetupTestData(appUrl+"/internal/status", `{}`, 200)

		err := Conformance(appUrl, httpFetcher, logger.NewMockLogger())

		Expect(err).To(Not(BeNil()))
	})

	It("Errors when non valid json returned for version", func() {
		appUrl := "https://app-under-test.cf.io"

		httpFetcher := NewMockHttpFetcher()
		httpFetcher.SetupTestData(appUrl+"/internal/version", `Wiiiie`, 200)
		httpFetcher.SetupTestData(appUrl+"/internal/status", `{}`, 200)

		err := Conformance(appUrl, httpFetcher, logger.NewMockLogger())

		Expect(err).To(Not(BeNil()))
	})

	It("Finds status", func() {
		appUrl := "https://app-under-test.cf.io"

		httpFetcher := NewMockHttpFetcher()
		httpFetcher.SetupTestData(appUrl+"/internal/version", `{"revision": "a840f1ae2122c3540a47541887eccff04aaeb212"}`, 200)
		httpFetcher.SetupTestData(appUrl+"/internal/status", `{}`, 200)

		err := Conformance(appUrl, httpFetcher, logger.NewMockLogger())

		Expect(err).To(BeNil())
	})

	It("Errors when non 200 for /internal/status", func() {
		appUrl := "https://app-under-test.cf.io"

		httpFetcher := NewMockHttpFetcher()
		httpFetcher.SetupTestData(appUrl+"/internal/version", `{"revision": "a840f1ae2122c3540a47541887eccff04aaeb212"}`, 200)
		httpFetcher.SetupTestData(appUrl+"/internal/status", `{}`, 404)

		err := Conformance(appUrl, httpFetcher, logger.NewMockLogger())

		Expect(err).To(Not(BeNil()))
	})
})
