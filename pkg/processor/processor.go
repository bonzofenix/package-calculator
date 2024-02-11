package processor

import (
	"fmt"
	"sort"
)

type Processor struct {
	packSizes []int
}

func NewProcessor() *Processor {
	return &Processor{}
}

func (p *Processor) AddPackSize(packSize int) {
	// Find where to insert element
	i := sort.Search(len(p.packSizes), func(i int) bool { return p.packSizes[i] <= packSize })
	// Insert  dummy element at the end of the slice to resize it
	p.packSizes = append(p.packSizes, 0)
	// Shift elements to the right
	copy(p.packSizes[i+1:], p.packSizes[i:])
	// Insert element in the right position
	p.packSizes[i] = packSize

}

func (p *Processor) GetPackSizes() []int {
	return p.packSizes
}

func (p *Processor) CalculatePacks(items int) [][]int {
	result := [][]int{}
	pendingItems := items

	// finds the minimum number of packs needed to fulfill the order
	for _, packSize := range p.packSizes {

		fmt.Println(packSize, pendingItems)

		if pendingItems/packSize > 0 {
			packCount := pendingItems / packSize
			result = append(result, []int{packSize, packCount})
			pendingItems = pendingItems % packSize
		}
	}
	return result
}
