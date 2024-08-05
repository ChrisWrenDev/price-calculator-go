package main

import (
  "example.com/price-calculator-go/prices"
)

func main(){
 var taxRates: []float{0,0.07,0.1,0.15}

 for _, taxRate := range taxRates {
   priceJob := prices.NewTaxIncludedPriceJob(taxRate)
   priceJob.Process()
 }

 fmt.Println(result)
}
