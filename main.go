package main

import (
	"github.com/Uno-count/Price-Calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.7, 0.1, 0, 0.15}

	for _, taxRates := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRates)
		priceJob.Process()
	}

}
