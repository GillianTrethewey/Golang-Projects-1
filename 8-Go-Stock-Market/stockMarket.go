package main

import (
	"fmt"
	"math/rand"
	"time"
)
func randomNumberGen(min float32, max float32) float32 {	
	return min + (max-min)*rand.Float32()
}
// Task implementation goes here
type Stock struct {
  name string
  price float32
}

func (s *Stock) updateStock() {
  change := randomNumberGen(-10000, 10000)
  s.price += change
}

func displayMarket(market []Stock) {
  for i := 0; i < len(market); i++ {
    fmt.Println(market[i])
  }
}

func main() {
  rand.Seed(time.Now().UnixNano())

  // GOOG := Stock{"GOOG", 2313.50}
  // AAPL := Stock{"AAPL", 157.28}
  // FB := Stock{"FB", 203.77}
  // TWTR := Stock{"TWTR", 50.00}
  //stockMarket := []Stock{GOOG, AAPL, FB, TWTR}
  stockMarket := []Stock{{"GOOG", 2313.50}, {"AAPL", 157.28}, {"FB", 203.77}, {"TWTR", 50.00}}
	// Function calls go here

displayMarket(stockMarket)
stockMarket[0].updateStock()
stockMarket[1].updateStock()


displayMarket(stockMarket)
}