package processor

import (
	"math"
	"sort"
)

// IProcessor enable injection of mock for testing
type IProcessor interface {
	CalculatePacks(packSizes []int, order int) map[int]int
}

// TODO: Remove this scruct and refactor code to just send the function around
// to avoid the need of a struct
type Processor struct {
}

func NewProcessor() *Processor {
	return &Processor{}
}

// Calculates the minimum number of packs required to fulfill an order,
// based on the available pack sizes:
// - Calculates the biggest order possible to fulfill with the largest pack size
// - Populates an array with the minimum number of packages required to fulfill different orders
// - Store the last pack size used to fullfill the order
// - Rebuild the result based on the smallest order possible with the last amout of pack sizes used
func (p *Processor) CalculatePacks(packs []int, order int) map[int]int {
	// when there is only one pack size, the order is fulfilled right away
	if len(packs) == 1 {
		return map[int]int{
			packs[0]: order/packs[0] + 1,
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(packs)))
	upperOrder := calculateUpperOrder(order, packs[0])

	packUsed := make([]int, upperOrder)
	minPackagesForItems := initializeMinPackages(upperOrder)

	populateMinPackages(order, upperOrder, packs, minPackagesForItems, packUsed)

	availableOrder := findClosestGreaterOrder(order, upperOrder, minPackagesForItems)
	result := buildResult(packs, availableOrder, packUsed)

	return result
}

// calculateUpperOrder calculates the upper order that is a multiple of the largest pack size.
// This result into the biggest permutation possible for an order which is fullfilling the order
// with the largest pack size.
func calculateUpperOrder(order, largestPack int) int {
	return (order/largestPack + 1) * largestPack
}

func initializeMinPackages(upperOrder int) []int {
	minPackagesForItems := make([]int, upperOrder)
	for i := range minPackagesForItems {
		minPackagesForItems[i] = math.MaxInt32
	}
	minPackagesForItems[0] = 0
	return minPackagesForItems
}

// populateMinPackages populates the minPackagesForItems array and the packUsed array
func populateMinPackages(order, upperOrder int, packs, minPackagesForItems, packUsed []int) {
	smallestPack := packs[len(packs)-1]
	for i := smallestPack; i < upperOrder; i++ {
		for _, pack := range packs {
			if i < pack {
				continue
			}
			minPackageCount := minPackagesForItems[i-pack]
			if minPackageCount != math.MaxInt32 && minPackageCount+1 < minPackagesForItems[i] {
				minPackagesForItems[i] = minPackageCount + 1
				packUsed[i] = pack
			}
		}
	}
}

// findClosestGreaterOrder finds the order with least amount of items possible
func findClosestGreaterOrder(order, upperOrder int, minPackagesForItems []int) int {
	closestGreaterOrder := order
	for i := order; i < upperOrder; i++ {
		if minPackagesForItems[i] < minPackagesForItems[closestGreaterOrder] {
			closestGreaterOrder = i
		}
	}
	return closestGreaterOrder
}

func buildResult(packs []int, closestGreaterOrder int, packUsed []int) map[int]int {
	result := make(map[int]int, len(packs))
	pendingOrder := closestGreaterOrder
	for pendingOrder > 0 {
		packSize := packUsed[pendingOrder]

		result[packSize]++
		pendingOrder -= packSize
	}
	return result
}
