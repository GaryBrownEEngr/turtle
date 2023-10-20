# Turtle

[![Go Reference](https://pkg.go.dev/badge/github.com/GaryBrownEEngr/turtle.svg)](https://pkg.go.dev/github.com/GaryBrownEEngr/turtle)
[![Go CI](https://github.com/GaryBrownEEngr/turtle/actions/workflows/go.yml/badge.svg)](https://github.com/GaryBrownEEngr/turtle/actions/workflows/go.yml)

## A Turtle Graphics System for Golang

Based on the python turtle, this Go package provides an environment to learn Go programming while getting instant feedback. The screen is updated in real time using the [Ebitengine](https://ebitengine.org/).

## Install

Go 1.20 or later is required.<br>
Ebitengine is the main dependency. [Check here the system specific instructions](https://ebitengine.org/en/documents/install.html).

## Example

### 5 Turtles At Once

```bash
go run github.com/GaryBrownEEngr/turtle/examples/turtlebasic@latest
```

![Example Picture](https://github.com/GaryBrownEEngr/turtle/blob/main/examples/turtlebasic/turtlebasic.png)

### Go Gopher

Converted from the python script seen in [this youtube video](https://www.youtube.com/watch?v=d8A1jqOGzNE).

```bash
go run github.com/GaryBrownEEngr/turtle/examples/fill@latest
```

![Example Picture](https://github.com/GaryBrownEEngr/turtle/blob/main/examples/fill/GoGopher.png)

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

	GoTo(x, y float64)      // Cartesian (x,y). Center in the middle of the window
	GetPos() (x, y float64) // Cartesian (x,y). Center in the middle of the window

	Left(angle float64)
	L(angle float64) // Turn Left alias
	Right(angle float64)
	R(angle float64) // Turn Right alias
	Angle(angle float64)
	GetAngle() float64
	PointToward(x, y float64)

	DegreesMode() // Default is degrees mode.
	RadiansMode()
	CompassMode() // Make it so North is 0 degrees, East is 90...
	GetAngleMode() AngleMode

	Speed(PixelsPerSecond float64)
	GetSpeed() float64

	PenUp()
	PU()  // Pen Up alias
	Off() // Pen Up alias
	PenDown()
	PD() // Pen Down alias
	On() // Pen Down alias
	Color(c color.RGBA)
	GetColor() color.RGBA
	Size(size float64)
	GetSize() float64
	Dot(size float64)
	Fill(c color.RGBA)

	Circle(radius, angleAmountToDraw float64, steps int)

	ShowTurtle()
	HideTurtle()    // Default
	ShapeAsTurtle() // Default
	ShapeAsArrow()
	ShapeAsImage(in image.Image)
	ShapeScale(scale float64)
}
```
