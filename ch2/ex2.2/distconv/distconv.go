package distconv

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }

// MToF はメートルをフィートに変換します
func MToF(m Meter) Feet { return Feet(3.28084 * m) }

// FToMはフィートをメートルに変換します
func FToM(f Feet) Meter { return Meter(0.3048 * f) }
