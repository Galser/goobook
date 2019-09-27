package lenconvert

import "fmt"

// Meters : length in meteres
type Meters float64

//Feet : length in imperial feet
type Feet float64

func (f Feet) String() string   { return fmt.Sprintf("%g ft", f) }
func (m Meters) String() string { return fmt.Sprintf("%g m", m) }

// FTToM - converts feet into meters
func FTToM(ft Feet) Meters { return Meters(ft * 0.3048) }

// MToFT - converts meteres into feet
func MToFT(m Meters) Feet { return Feet(m * 3.280839895) }
