package tempconv

import (
	"fmt"
)

//Celsius
type Celsius float64

//Fahreneit
type Fahreneit float64

//Kelvin
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string   { return fmt.Sprintf("%g°C", c) }
func (f Fahreneit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string    { return fmt.Sprintf("%g°K", k) }
