// ############## Page 44 ##############
package main

import "fmt"

/*
  Gear
*/
type Gear struct {
	Chainring, Cog float64
	Wheel          Diameter
}

type Diameter interface {
	Diameter() float64
}

func NewGear(chainring, cog float64, wheel Diameter) *Gear {
	return &Gear{chainring, cog, wheel}
}

func (gear Gear) GearInches() float64 {
	// a few lines of scary math
	return gear.Ratio() * gear.diameter()
	// more lines of scary math
}

/*
  isolate external messages (if it may change, or is used from several places,
  or within complex code)
*/
func (gear Gear) diameter() float64 {
	return gear.Wheel.Diameter()
}

func (gear Gear) Ratio() float64 {
	return gear.Chainring / gear.Cog
}

/*
  Wheel
*/
type Wheel struct {
	Rim, Tire float64
}

func NewWheel(rim, tire float64) *Wheel {
	return &Wheel{rim, tire}
}

func (wheel Wheel) Diameter() float64 {
	return wheel.Rim + (wheel.Tire * 2)
}

/*
  Main
*/
func main() {
	gear := NewGear(52, 11, NewWheel(26, 1.5))
	fmt.Println(gear.GearInches()) // => 137.0909090909091
	fmt.Println(gear.Ratio())      // => 4.7272727272727275
}
