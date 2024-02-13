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
	CalculatePacksCalledWith []int
	AddPackSizesCalledWith   []int
}

func (m *mockProcessor) CalculatePacks(packSizes []int, order int) map[int]int {

	m.CalculatePacksCalledWith = append(m.CalculatePacksCalledWith, order)
	return map[int]int{}
}

func (m *mockProcessor) AddPackSize(packSize int) {
	m.AddPackSizesCalledWith = append(m.AddPackSizesCalledWith, packSize)
}

func (m *mockProcessor) GetPackSizes() []int {
	return []int{}
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

		BeforeEach(func() {
			fakeProcessor = &mockProcessor{}

			server.AppendHandlers(
				CalculateHandler(fakeProcessor),
			)
		})

		It("it should re render index with response", func() {
			resp, err := http.Post(server.URL()+"/calculate", "application/x-www-form-urlencoded", nil)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))
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
			Expect(fakeProcessor.CalculatePacksCalledWith[0]).Should(Equal(300))
		})
	})
})
