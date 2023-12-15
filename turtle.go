package turtle

import (
	"github.com/GaryBrownEEngr/turtle/ebitencanvas"
	"github.com/GaryBrownEEngr/turtle/pen"
	"github.com/GaryBrownEEngr/turtle/turtlemodel"
)

// Top level interface that most users of the turtle package will exclusively use.
type Window interface {
	GetCanvas() turtlemodel.Canvas
	NewTurtle() turtlemodel.Turtle
	NewTurtlePermissive() turtlemodel.TurtlePermissive
}

type window struct {
	can turtlemodel.Canvas
}

var _ Window = &window{}

// Get the canvas interface
func (s *window) GetCanvas() turtlemodel.Canvas {
	return s.can
}

// Create a new turtle
func (s *window) NewTurtle() turtlemodel.Turtle {
	return pen.NewPen(s.can)
}

// Create a new turtle that converts all input number from any real number type to float64.
func (s *window) NewTurtlePermissive() turtlemodel.TurtlePermissive {
	return pen.NewPenPermissive(s.can)
}

///////////////////////////////////////////////////
///////////////////////////////////////////////////

// Turtle Window creation parameters
type Params struct {
	Width   int  // pixels
	Height  int  // pixels
	ShowFPS bool // Show frame per second and other info in the top left corner.
}

// Wrap the starting of ebitencanvas inside something that implements the Window interface
// so that most of the time a user will only need one import statement from this repo to make a turtle graphic.
// But the actual game, drawing, and sprite implementations can still be separated nicely into packages.
func Start(params Params, drawFunc func(Window)) {
	canvasParams := ebitencanvas.CanvasParams{
		Width:   params.Width,
		Height:  params.Height,
		ShowFPS: params.ShowFPS,
	}

	// Create a callback that translates the models.Canvas into a Window
	initCallback := func(can turtlemodel.Canvas) {
		drawFunc(&window{can: can})
	}
	ebitencanvas.StartEbitenTurtleCanvas(canvasParams, initCallback)
}

// Get the newest keyboard/mouse just pressed event from the given channel.
// This returns nil if there is no new data.
// This will throw away all but the newest set of data available. So this should be called faster that the game update rate (60Hz),
// otherwise sim.PressedUserInput() should be used instead.
func GetNewestJustPressedFromChan(justPressedChan chan *turtlemodel.UserInput) *turtlemodel.UserInput {
	var ret *turtlemodel.UserInput

ChanExtractionLoop:
	for {
		select {
		case i := <-justPressedChan:
			ret = i
		default:
			// receiving from chan would block
			break ChanExtractionLoop
		}
	}
	return ret
}
