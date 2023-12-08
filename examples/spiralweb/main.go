package main

import (
	"image/color"

	"github.com/GaryBrownEEngr/turtle"
)

func main() {
	params := turtle.Params{Width: 500, Height: 500}
	turtle.Start(params, drawFunc)
}

// drawFunc is started as a goroutine.
// https://www.geeksforgeeks.org/draw-colorful-spiral-web-using-turtle-graphics-in-python/
func drawFunc(window turtle.Window) {
	// go turtle.CreateGif(window, time.Millisecond*400, time.Millisecond*150, "./examples/spiralweb/spiralweb.gif", 60)

	window.GetCanvas().ClearScreen(turtle.Black)
	t := window.NewTurtle()
	t.ShowTurtle()
	t.Speed(1000)
	t.PenDown()

	colors := []color.Color{turtle.Red, turtle.Yellow, turtle.Green, turtle.Purple, turtle.Blue, turtle.Orange}
	for x := 1; x < 200; x++ {
		t.Color(colors[x%6])        // setting color
		t.Size(float64(x)/75.0 + 1) // setting width
		t.Forward(float64(x))       // moving forward
		t.Left(59)                  // Turn left
	}
}
