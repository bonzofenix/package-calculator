package app_test

import (
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/bonzofenix/package-calculator/app"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Server", func() {
	var server *ghttp.Server

	BeforeEach(func() {
		server = ghttp.NewServer()
	})

	AfterEach(func() {
		server.Close()
	})

	Context("When get request is sent to empty path", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				RootHandler,
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
})
