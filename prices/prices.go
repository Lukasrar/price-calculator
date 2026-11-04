package prices

import (
	"fmt"
	iomanager "lukasrar/price-calculator/IOManager"
	"lukasrar/price-calculator/conversion"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, readFileError := job.IOManager.ReadLines()

	if readFileError != nil {
		return readFileError
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
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

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {

	err := job.LoadData()

	if err != nil {
		// return err

		errorChan <- err
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
