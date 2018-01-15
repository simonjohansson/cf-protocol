package in_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/simonjohansson/cf-protocol/resource/in"
)

var _ = Describe("In", func() {
	Context("when the version is given on input", func() {
		It("should output the version that it was given", func() {
			inputData := `
	{
	  "source": {
	    "a-key": "some data",
	    "another-key": "some more data"
	  },
	  "version": { "ref": "61cebf" }
	}
				`
			output, err := in.Execute([]byte(inputData))
			Ω(err).ShouldNot(HaveOccurred())
			Expect(output).To(MatchJSON(`{"version": { "ref": "61cebf" }}`))
		})
	})

	Context("when the version is not given on input", func() {
		It("should return an error", func() {
			output, err := in.Execute([]byte(`{ "missing" : "the version" }`))
			Ω(err).Should(HaveOccurred())
			Expect(err.Error()).Should(Equal("missing version"))
			Expect(output).To(BeEmpty())
		})
	})

	Context("when bad json given on input", func() {
		It("should return an error", func() {
			output, err := in.Execute([]byte(``))
			Ω(err).Should(HaveOccurred())
			Expect(err.Error()).Should(Equal("unexpected end of JSON input"))
			Expect(output).To(BeEmpty())
		})
	})

})
