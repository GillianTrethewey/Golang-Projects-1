package main

import "fmt"
import "math"



func main() {
  currencies := map[string]float64{
    "JPY": 130.2,
    "EUR": 0.95,
    "CDN": 1.36,
    "GBP": 0.92,
  }

  var dollarAmount float64
  var currency string

  fmt.Println("Enter the dollar amount:")
  fmt.Scan(&dollarAmount)
  if (dollarAmount == 0 ) {
    fmt.Println("Please enter a valid dollar value")
  } else {
    fmt.Println("You entered $", dollarAmount)
  }

  fmt.Println("Enter the target currency:")
  fmt.Scan(&currency)

  rate, isValid := currencies[currency]
  if(isValid) {
    fmt.Println("The converted amount is: $", math.Round(rate*dollarAmount*100)/100)
  } else {
    fmt.Println(currency, "was not found")
  }
  

}