package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
import (
	"math"
)

type Num int

const SidesTriangle, SidesSquare, SidesCircle = 3, 4, 0

func CalcSquare(sideLen float64, sidesNum Num) float64 {
	var n float64
	if sidesNum == SidesSquare {
		n = sideLen * sideLen
	} else if sidesNum == SidesTriangle {
		n = math.Sqrt(3) / 4 * math.Pow(sideLen, 2)
	} else if sidesNum == SidesCircle {
		n = math.Pi * math.Pow(sideLen, 2)
	} else {
		n = 0
	}
	return n
}
