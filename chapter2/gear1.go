// ############## Page 19 ##############
package main

import "fmt"

type Gear struct {
  Chainring int // number of teath
  Cog       int
}

func NewGear(chainring, cog int) Gear {
  return Gear{chainring, cog}
}

func (gear Gear) Ratio() float64 {
  return float64(gear.Chainring) / float64(gear.Cog)
}

func main() {
  fmt.Println(NewGear(52, 11).Ratio()) // => 4.7272727272727275
  fmt.Println(NewGear(30, 27).Ratio()) // => 1.1111111111111112
}
