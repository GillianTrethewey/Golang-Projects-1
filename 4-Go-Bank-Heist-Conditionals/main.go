package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  var eludedGuards int
  eludedGuards = rand.Intn(100)

  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }

  openedVault := rand.Intn(100)

  if isHeistOn == true && openedVault >= 20 {
    fmt.Println("Grab and GO!")
  } else if isHeistOn == true && openedVault < 70 {
    isHeistOn = false
    fmt.Println("The vault can't be opened.")
  }

  leftSafely := rand.Intn(5)

  if isHeistOn {
    switch leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("Looks like you tripped an alarm... run?")
      case 1:
        isHeistOn = false
        fmt.Println("Turns out value doors don't open from the inside...")
      case 2: 
        isHeistOn = false
        fmt.Println("Looks like this fingerprint scanner won't accept any fingerprint...")
      case 3:
        isHeistOn = false
        fmt.Println("Did I even pack the burlap bags?")
      case 4:
        isHeistOn = true
        fmt.Println("You did it!")
      default:
        fmt.Println("Start the getaway car!")
    }
      
  }

  if isHeistOn {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Println("$", amtStolen, "not bad!")
  }

  fmt.Println("isHeistOn is currently:", isHeistOn)
}

