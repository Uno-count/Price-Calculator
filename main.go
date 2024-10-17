package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.1, 0, 0.15}

	result := make(map[float64][]float64)

	for _, taxRates := range taxRates {
		taxIncludedPrice := make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludedPrice[priceIndex] = price * (1 + taxRates)
		}
		result[taxRates] = taxIncludedPrice
	}

	fmt.Println(result)
}
