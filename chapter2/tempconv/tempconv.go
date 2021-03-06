// Package tempconv erforms Celsius and Fahrenheit conversions
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100.
)

func (c Celsius) String() string {
	return fmt.Sprintf("%f°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%f°F", f)
}
