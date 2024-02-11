package processor_test

import (
	"github.com/bonzofenix/package-calculator/pkg/processor"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Processor", func() {
	Describe("NewProcessor", func() {
		It("should return a new processor", func() {
			p := processor.NewProcessor()
			Expect(p).NotTo(BeNil())
		})
	})

	Describe("AddPackSize", func() {
		It("should add a package size to the processor", func() {
			p := processor.NewProcessor()
			p.AddPackSize(11)
			Expect(p.GetPackSizes()).To(ContainElement(11))
		})

		It("should add package in order", func() {
			p := processor.NewProcessor()
			p.AddPackSize(3)
			p.AddPackSize(1)
			p.AddPackSize(2)
			Expect(p.GetPackSizes()).To(Equal([]int{1, 2, 3}))
		})
	})

	Describe("CalculatePacks", func() {
		It("should return the number of packs for a given size", func() {
			p := processor.NewProcessor()
			p.AddPackSize(23)
			p.AddPackSize(31)
			p.AddPackSize(53)
			Expect(p.CalculatePacks(263)).To(Equal([][]int{
				[]int{2, 23},
				[]int{7, 31},
			}))
		})
	})
})
