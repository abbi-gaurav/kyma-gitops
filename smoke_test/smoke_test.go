package smoke_test

import (
	"fmt"
	"github.com/onsi/ginkgo/reporters"
	"net/http"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter(fmt.Sprintf("/tmp/reports/%s_junit.xml", os.Getenv("TEST_NAME")))
	RunSpecsWithDefaultAndCustomReporters(t, "API Suite", []Reporter{junitReporter})

}

var httpClient *http.Client
var apiURL string

var _ = BeforeSuite(func() {
	httpClient = &http.Client{}
	apiURL = os.Getenv("API_URL")
})

var _ = Describe("My Kyma API test suite", func() {
	Describe("GET /", func() {
		Context("When GET method is called", func() {
			It(fmt.Sprintf("should return [%s]", "hello world"), func() {
				req, _ := http.NewRequest(http.MethodGet, apiURL, nil)
				resp, err := httpClient.Do(req)
				Expect(err).Should(BeNil())
				Expect(resp).ShouldNot(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
			})
		})
	})
})
