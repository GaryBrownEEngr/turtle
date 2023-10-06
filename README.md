# Turtle

[![Go Reference](https://pkg.go.dev/badge/github.com/GaryBrownEEngr/turtle.svg)](https://pkg.go.dev/github.com/GaryBrownEEngr/turtle)

## A Turtle Graphics System for Golang

Based on the python turtle, this Go package provides an environment to learn Go programming while getting instant feedback. The screen is updated in real time using the [Ebitengine](https://ebitengine.org/).

## Install

Ebitengine is the main dependency. [Check here the system specific instructions](https://ebitengine.org/en/documents/install.html).

## Example

```bash
go run github.com/GaryBrownEEngr/turtle/examples/turtlebasic@latest
```

![Example Picture](https://github.com/GaryBrownEEngr/turtle/blob/main/examples/turtlebasic/turtlebasic.png)

## Basic Example Program

```go
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
	t.SetDegreesMode()
	t.PenColor(turtleutil.White)
	t.SetSpeed(1000)
	t.PenDown()

	t.Left(45)
	t.Forward(100)
	t.Left(135)
	t.Forward(200 * 0.707)
	t.GoTo(0, 0)
	t.SetAngle(-15)
	for i := 0; i < 12; i++ {
		t.Forward(100)
		t.Right(30)
	}
}
```

## Turtle Controls

When a new turtle is created, it is given a Canvas interface. The turtle itself fulfills the Turtle interface seen below.
Which means each turtle created can perform this exact list of actions.

```go
type Turtle interface {
	Forward(distance float64)
	F(distance float64) // Forward alias
	Backward(distance float64)
	B(distance float64) // Backward alias
	PanRightward(distance float64)
	PanR(distance float64) // PanRightward alias
	PanLeftward(distance float64)
	PanL(distance float64) // PanLeftward alias

	GoTo(x, y float64)
	GetPos() (x, y float64)

	Left(angle float64)
	L(angle float64) // Turn Left alias
	Right(angle float64)
	R(angle float64) // Turn Right alias
	SetAngle(angle float64)
	PointToward(x, y float64)
	GetAngle() float64

	SetDegreesMode() // Default is degree mode.
	SetRadianMode()
	EnableCompassAngleMode(in bool) // Make it so North is 0 degrees, East is 90...

	SetSpeed(PixelsPerSecond float64)

	PenUp()
	PU()  // Pen Up alias
	Off() // Pen Up alias
	PenDown()
	PD() // Pen Down alias
	On() // Pen Down alias
	PenColor(c color.RGBA)
	PenSize(size float64)
}
```
