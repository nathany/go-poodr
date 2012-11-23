// ############## Page 20 ##############
package main

import "fmt"

type Gear struct {
  Chainring float64 // number of teath
  Cog       float64
  Rim       float64 // diameters
  Tire      float64
}

func NewGear(chainring, cog, rim, tire float64) Gear {
  return Gear{chainring, cog, rim, tire}
}

func (gear Gear) Ratio() float64 {
  return gear.Chainring / gear.Cog
}

func (gear Gear) GearInches() float64 {
  return gear.Ratio() * (gear.Rim + (gear.Tire * 2))
}

func main() {
  fmt.Println(NewGear(52, 11, 26, 1.5).GearInches())  // => 137.0909090909091
  fmt.Println(NewGear(52, 11, 24, 1.25).GearInches()) // => 125.27272727272728
}
