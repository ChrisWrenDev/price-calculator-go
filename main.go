package main

import (
  "fmt"
	"example.com/price-calculator/prices"
  "example.com/price-calculator/filemanager"
  "example.com/price-calculator/cmdmanager"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
  doneChans := make([]chan bool, len(taxRates))
  errChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
    doneChans[index] = make(chan bool)
    errChans[index] = make(chan error)

    fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))
    cmd := cmdmanager.New() 

    priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
    go priceJob.Process(doneChans[index], errChans[index])

 	}

  for index  := range taxRates {
    select {
      case err := <- errChans[index]:
        if err != nil {
        fmt.Println(err)
      }
    case <- doneChans[index]:
      fmt.Println("Done!")
    }
  }
}
