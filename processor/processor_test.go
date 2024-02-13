package processor_test

import (
	"github.com/bonzofenix/package-calculator/processor"

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

	Describe("CalculatePacks", func() {
		It("return should work for single packages", func() {
			p := processor.NewProcessor()
			Expect(p.CalculatePacks([]int{10}, 24)).To(Equal(map[int]int{10: 3}))
		})

		It("return 1 pack when items match available sizes", func() {
			p := processor.NewProcessor()

			Expect(p.CalculatePacks([]int{10, 20}, 20)).To(Equal(map[int]int{20: 1}))
		})

		It("should return the number of packs for a given size", func() {
			p := processor.NewProcessor()
			Expect(p.CalculatePacks([]int{23, 53, 31}, 263)).To(Equal(map[int]int{23: 2, 31: 7}))
		})
	})
})
