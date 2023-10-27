package main

import (
	"math"
	"time"

	"github.com/GaryBrownEEngr/turtle"
	"github.com/GaryBrownEEngr/turtle/models"
)

// Try to touch the worm. It runs away. The closer you get the faster it goes. But if you catch it, it explodes.
// Press c to clear the screen.
func main() {
	params := turtle.Params{Width: 1000, Height: 1000, ShowFPS: true}
	turtle.Start(params, drawFunc)
}

func drawFunc(window turtle.Window) {
	can := window.GetCanvas()
	can.ClearScreen(turtle.White)
	t := window.NewTurtle()
	t.ShowTurtle()
	t.PenDown()
	t.Size(5)
	speed := 300.0

	prevUserIn := &models.UserInput{}
	for {
		userIn := can.GetUserInput()
		if userIn.KeysDown.C && !prevUserIn.KeysDown.C {
			can.ClearScreen(turtle.White)
		}

		tX, tY := t.GetPos()
		mX, mY := float64(userIn.MouseX), float64(userIn.MouseY)

		t.PointToward(mX, mY)
		t.Right(180)

		detlaX := tX - mX
		deltaY := tY - mY
		dist := math.Sqrt(detlaX*detlaX + deltaY*deltaY)

		if dist < 300 {
			ratio := 2.0 * (1.0 - dist/100.0)
			switch {
			case ratio <= 1:
				t.Color(turtle.LerpColor(turtle.Blue, turtle.Green, ratio))
			case ratio > 1:
				t.Color(turtle.LerpColor(turtle.Green, turtle.Red, ratio-1.0))
			}

			if dist < 5 {
				t.Dot(100)
			} else {
				speed := 9.0 * 300.0 / dist
				t.Speed(speed)
			}
			t.Forward(speed * 0.010)
		}

		prevUserIn = &userIn
		time.Sleep(10 * time.Millisecond)
	}
}
