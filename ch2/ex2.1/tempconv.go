package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
	BoilingK      Kelvin  = 273.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

// CToF は摂氏の温度を華氏へ変換します。
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC は華氏の温度を摂氏へ変換します。
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

func FToK(f Fahrenheit) Kelvin { return Kelvin(CToK(FToC(f))) }

func KToC(k Kelvin) Celsius { return Celsius(k + CToK(AbsoluteZeroC)) }

func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }
