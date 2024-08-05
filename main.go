package main

func main(){
 prices := []float{10,20,30}
 var taxRates: []float{0,0.07,0.1,0.15}

 var result map[float64][]float64

 for index, taxRate := range taxRates {
   var taxIncludedPrices []float64 = make([]float64, len(prices))
   for priceIndex, price := prices {
     taxIncludedPrices[priceIndex] = price * (1 + taxRate)
   }
   result[taxRate] = taxIncludedPrices
 }

 fmt.Println(result)
}
