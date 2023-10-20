package main

import (
	"time"

	"github.com/GaryBrownEEngr/turtle/ebitencanvas"
	"github.com/GaryBrownEEngr/turtle/models"
	"github.com/GaryBrownEEngr/turtle/turtle"
	"github.com/GaryBrownEEngr/turtle/turtleutil"
)

// Drag the mouse with the left button pressed to draw on the canvas.
// Press c to clear the screen.
func main() {
	params := ebitencanvas.CanvasParams{Width: 1000, Height: 1000, ShowFPS: true}
	ebitencanvas.StartEbitenTurtleCanvas(params, drawFunc)
}

func drawFunc(can models.Canvas) {
	can.ClearScreen(turtleutil.White)
	var t models.Turtle = turtle.NewTurtle(can)
	t.Speed(1e9)

	ratio := 0.0
	prevUserIn := &models.UserInput{}
	for {
		userIn := can.GetUserInput()
		if userIn.KeysDown.C && !prevUserIn.KeysDown.C {
			can.ClearScreen(turtleutil.White)
		}

		if userIn.MouseDown.Left {
			desiredColor := turtleutil.LerpColor(turtleutil.Blue, turtleutil.Red, ratio)
			t.Color(desiredColor)
			t.GoTo(float64(prevUserIn.MouseX), float64(prevUserIn.MouseY))
			t.PenDown()
			t.GoTo(float64(userIn.MouseX), float64(userIn.MouseY))
			t.PenUp()
		}
		prevUserIn = &userIn
		ratio += .001
		if ratio > 1 {
			ratio = 0
		}
		time.Sleep(1 * time.Millisecond)
	}
}
