//  My homage to 10 print chr$(205.5+rnd(1)) goto 10
// A famous c64 program that prints out a maze.
// Example here: https://www.youtube.com/watch?v=m9joBLOZVEo
// By Marc Paquette 
// November 2017

package main

import (
  "fmt"
  "math/rand"
  )

func main () {
  rand.Seed(42)
  walls := []string{
    "╲",
    "╱",
  }

  for {
    fmt.Printf("%s",walls[rand.Intn(len(walls))])
  }
}
