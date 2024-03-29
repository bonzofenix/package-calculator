package app_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/bonzofenix/package-calculator/app"
	"github.com/onsi/gomega/ghttp"
)

type mockProcessor struct {
	CalculatePacksCalledWith []calculatePacksArguments
}

type calculatePacksArguments struct {
	PackSizes []int
	Order     int
}

func (m *mockProcessor) CalculatePacks(packSizes []int, order int) map[int]int {
	arguments := calculatePacksArguments{packSizes, order}

	m.CalculatePacksCalledWith = append(m.CalculatePacksCalledWith, arguments)

	return map[int]int{}
}

var _ = Describe("Server", func() {
	var server *ghttp.Server

	BeforeEach(func() {
		server = ghttp.NewServer()
	})

	AfterEach(func() {
		server.Close()
	})

	Context("Get /", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				RootHandler(),
			)
		})

		It("Returns the empty path", func() {
			resp, err := http.Get(server.URL() + "/")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))
			body, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(body)).To(ContainSubstring("PACKAGE CALCULATOR"))
		})
	})

	Context("POST /calculate", func() {
		var fakeProcessor *mockProcessor
		var urlPath string

		BeforeEach(func() {
			urlPath = server.URL() + "/calculate"
			fakeProcessor = &mockProcessor{}

			server.AppendHandlers(
				CalculateHandler(fakeProcessor),
			)
		})

		It("it should not fail when no data is submited", func() {
			formData := url.Values{}
			formEncoded := formData.Encode()
			resp, err := http.Post(urlPath, "application/x-www-form-urlencoded", bytes.NewBufferString(formEncoded))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusBadRequest))
		})

		It("it should send the data from the form to the processor", func() {
			formData := url.Values{
				"packSizes": {"100,200,300"},
				"order":     {"300"},
			}

			formEncoded := formData.Encode()
			Expect(fakeProcessor.CalculatePacksCalledWith).Should(HaveLen(0))
			resp, err := http.Post(server.URL()+"/calculate", "application/x-www-form-urlencoded", bytes.NewBufferString(formEncoded))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))
			Expect(fakeProcessor.CalculatePacksCalledWith).Should(HaveLen(1))

			Expect(fakeProcessor.CalculatePacksCalledWith[0].PackSizes).Should(Equal([]int{100, 200, 300}))
			Expect(fakeProcessor.CalculatePacksCalledWith[0].Order).Should(Equal(300))
		})
	})
})
