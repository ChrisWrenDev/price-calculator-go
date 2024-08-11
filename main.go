package main

import (
  "fmt"
	"example.com/price-calculator/prices"
  "example.com/price-calculator/filemanager"
  "example.com/price-calculator/cmdmanager"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
    fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))
    cmd := cmdmanager.New() 

    priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
    err := priceJob.Process()

    if err != nil {
      fmt.Print("Could not process job")
      fmt.Print(err)
    }
	}

}
