package turtle

import (
	"image/color"

	"github.com/GaryBrownEEngr/turtle/turtleutil"
	"golang.org/x/exp/constraints"
)

// Provide a pass through layer so users of this library don't have to import the tools package.

// Linearly interpolate between any two numbers of the same type using a ratio.
// The ratio is capped between 0 and 1
func Lerp[T constraints.Integer | constraints.Float](a, b T, ratio float64) T {
	return turtleutil.Lerp(a, b, ratio)
}

// Creates a color between the given a and b. 0 means a is given, 1 means b is given, .5 is a color half way between.
// The ratio is capped between 0 and 1
// Currently the function floors the number instead of rounding to nearest.
func LerpColor(a, b color.RGBA, ratio float64) color.RGBA {
	return turtleutil.LerpColor(a, b, ratio)
}
