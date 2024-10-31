package main

import (
	"github.com/Uno-count/Price-Calculator/cmd"
	"github.com/Uno-count/Price-Calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.7, 0.1, 0, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices/data.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))

		cmdm := cmd.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob.Process()
	}

}
