package main

import (
	"fmt"

	"github.com/Uno-count/Price-Calculator/filemanager"
	"github.com/Uno-count/Price-Calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.7, 0.1, 0, 0.15}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)

		fm := filemanager.New("prices/data.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmd.New()

		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[i], errorChans[i])

		// if err != nil {
		// 	fmt.Println("cannot process the job")
		// 	fmt.Println(err)
		// }
	}

	for i, _ := range taxRates {
		select {
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Println(err)
			}

		case <-doneChans[i]:
			fmt.Println("Done!!!")
		}
	}

}
