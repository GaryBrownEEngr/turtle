package main

import (
	"github.com/GaryBrownEEngr/turtle/ebitencanvas"
	"github.com/GaryBrownEEngr/turtle/models"
	"github.com/GaryBrownEEngr/turtle/turtle"
	"github.com/GaryBrownEEngr/turtle/turtleutil"
)

func main() {
	// Create the Ebitengine canvas.
	params := ebitencanvas.CanvasParams{Width: 1000, Height: 1000}
	ebitencanvas.StartEbitenTurtleCanvas(params, drawFunc)
}

// drawFunc is started as a goroutine.
func drawFunc(can models.Canvas) {
	var t models.Turtle = turtle.NewTurtle(can)
	t.Color(turtleutil.White)
	t.Speed(1000)
	t.PenDown()

	t.Left(45)
	t.Forward(100)
	t.Left(135)
	t.Forward(200 * 0.707)
	t.GoTo(0, 0)
	t.Angle(-15)
	for i := 0; i < 12; i++ {
		t.Forward(100)
		t.Right(30)
	}
}
