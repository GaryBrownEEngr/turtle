package turtle

import (
	"image"
	"image/color"

	"github.com/GaryBrownEEngr/turtle/tools"
)

// Provide a pass through layer to users of this library don't have to import the tools package.

// The ratio is capped between 0 and 1
func Lerp(a, b, ratio float64) float64 {
	return tools.Lerp(a, b, ratio)
}

// The ratio is capped between 0 and 1
// Currently the function floors the number instead of rounding to nearest.
func LerpUint8(a, b uint8, ratio float64) uint8 {
	return tools.LerpUint8(a, b, ratio)
}

// Creates a color between the given a and b. 0 means a is given, 1 means b is given, .5 is a color half way between.
// The ratio is capped between 0 and 1
// Currently the function floors the number instead of rounding to nearest.
func LerpColor(a, b color.RGBA, ratio float64) color.RGBA {
	return tools.LerpColor(a, b, ratio)
}

func LoadSpriteFile(path string) (image.Image, error) {
	return tools.LoadSpriteFile(path)
}
