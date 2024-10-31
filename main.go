package main

import (
	"fmt"

	"github.com/Uno-count/Price-Calculator/filemanager"
	"github.com/Uno-count/Price-Calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.7, 0.1, 0, 0.15}

	doneChans := make([]chan bool, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		fm := filemanager.New("prices/data.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmd.New()

		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[i])

		// if err != nil {
		// 	fmt.Println("cannot process the job")
		// 	fmt.Println(err)
		// }
	}

	for _, doneChan := range doneChans {
		<-doneChan
	}

}
