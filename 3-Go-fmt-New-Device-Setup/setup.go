package main

import (
  "fmt"
)

func main() {
  var name string
  fmt.Println("What is your name?")
  fmt.Scan(&name)

  fmt.Printf("Hello %s \n", name)

  var age int
  fmt.Println("What is your age?")
  fmt.Scan(&age)

  fmt.Printf("Your name is %s, and your age is %d. \n", name, age)

  newMessage := fmt.Sprintf("Your name is really %s, and your age is really %d.\n", name, age)
  fmt.Print(newMessage)


// run using go run setup.go