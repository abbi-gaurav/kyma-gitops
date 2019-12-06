package smoke_test

import (
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/onsi/ginkgo/reporters"
	"net/http"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("/tmp/reports/junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "API Suite", []Reporter{junitReporter})

}

var httpClient *http.Client

var _ = BeforeSuite(func() {
	httpClient = &http.Client{}
})

var _ = Describe("My Kyma API ",
	func() {
		DescribeTable("Test suite",
			func(url string) {
				req, _ := http.NewRequest(http.MethodGet, url, nil)
				resp, err := httpClient.Do(req)
				Expect(err).Should(BeNil())
				Expect(resp).ShouldNot(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
			},
			Entry("Lambda API URL", os.Getenv("LAMBDA_API_URL")),
		)
	})
