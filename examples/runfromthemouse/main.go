package main

import (
	"math"
	"time"

	"github.com/GaryBrownEEngr/turtle/ebitencanvas"
	"github.com/GaryBrownEEngr/turtle/models"
	"github.com/GaryBrownEEngr/turtle/turtle"
	"github.com/GaryBrownEEngr/turtle/turtleutil"
)

// Try to touch the worm. It runs away. The closer you get the faster it goes. But if you catch it, it explodes.
// Press c to clear the screen.
func main() {
	params := ebitencanvas.CanvasParams{Width: 1000, Height: 1000, ShowFPS: true}
	ebitencanvas.StartEbitenTurtleCanvas(params, drawFunc)
}

func drawFunc(can models.Canvas) {
	can.ClearScreen(turtleutil.White)
	var t models.Turtle = turtle.NewTurtle(can)
	t.PenDown()
	t.PenSize(5)
	speed := 300.0

	prevUserIn := &models.UserInput{}
	for {
		userIn := can.GetUserInput()
		if userIn.KeysDown.C && !prevUserIn.KeysDown.C {
			can.ClearScreen(turtleutil.White)
		}

		tX, tY := t.GetPos()
		mX, mY := float64(userIn.MouseX), float64(userIn.MouseY)

		t.PointToward(mX, mY)

		detlaX := tX - mX
		deltaY := tY - mY
		dist := math.Sqrt(detlaX*detlaX + deltaY*deltaY)

		if dist < 300 {
			ratio := 2.0 * (1.0 - dist/300.0)
			switch {
			case ratio <= 1:
				t.PenColor(turtleutil.LerpColor(turtleutil.Blue, turtleutil.Green, ratio))
			case ratio > 1:
				t.PenColor(turtleutil.LerpColor(turtleutil.Green, turtleutil.Red, ratio-1.0))
			}

			if dist < 5 {
				t.PaintDot(100)
			} else {
				speed := 9.0 * 300.0 / dist
				t.SetSpeed(speed)
			}
			t.Backward(speed * 0.010)
		}

		prevUserIn = &userIn
		time.Sleep(10 * time.Millisecond)
	}
}
