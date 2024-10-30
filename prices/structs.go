package prices

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

// func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
// 	return &TaxIncludedPriceJob{
// 		InputPrices: []float64{10, 20, 30},
// 		TaxRate:     taxRate,
// 	}
// }
