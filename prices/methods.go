package prices

import (
	"fmt"

	"github.com/Uno-count/Price-Calculator/conversion"
	"github.com/Uno-count/Price-Calculator/iomanager"
)

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, chanError chan error) {

	err := job.LoadData()

	if err != nil {
		chanError <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)

	doneChan <- true

}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)

		return err
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return err
	}

	job.InputPrices = prices

	return nil
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
