package ebitencanvas

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/GaryBrownEEngr/turtle/models"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

/*
The base Ebiten "Game" is based off the example https://ebitengine.org/en/examples/noise.html
*/

type game struct {
	screenWidth  int
	screenHeight int
	showFPS      bool
	commands     chan drawCmd
	img          *image.RGBA
	controlState SavedControlState
	controls     *models.UserInput
}

func newGame(width, height int, showFPS bool, commands chan drawCmd) *game {
	g := &game{
		screenWidth:  width,
		screenHeight: height,
		showFPS:      showFPS,
		commands:     commands,
		img:          image.NewRGBA(image.Rect(0, 0, width, height)),
	}

	// ebiten.SetTPS(120)
	// ebiten.SetVsyncEnabled(false) // For some reason, on Windows, there is quite a bit of lag.
	// setting this to false clears it up, but also makes it run at 1000Hz...
	ebiten.SetWindowSize(g.screenWidth, g.screenHeight)
	ebiten.SetWindowTitle("Go Turtle Graphics")
	return g
}

// This function will not return. It must be run on the main thread.
func (g *game) runGame() {
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func clearImage(i *image.RGBA, c color.RGBA) {
	xMin := i.Rect.Min.X
	yMin := i.Rect.Min.Y
	xMax := i.Rect.Max.X
	yMax := i.Rect.Max.Y
	for y := yMin; y < yMax; y++ {
		for x := xMin; x < xMax; x++ {
			i.SetRGBA(x, y, c)
		}
	}
}

func bucketFill(i *image.RGBA, x, y int, c color.RGBA) {
	colorMatches := func(a, b color.RGBA) bool {
		return a.R == b.R && a.G == b.G && a.B == b.B && a.A == b.A
	}

	type upNextStruct struct {
		x, y int
	}
	upNextStack := []upNextStruct{}

	xMin := i.Rect.Min.X
	yMin := i.Rect.Min.Y
	xMax := i.Rect.Max.X
	yMax := i.Rect.Max.Y

	srcColor := i.RGBAAt(x, y)
	upNextStack = append(upNextStack, upNextStruct{x: x, y: y})

	for len(upNextStack) > 0 {
		xy := upNextStack[len(upNextStack)-1]
		x := xy.x
		y := xy.y
		upNextStack = upNextStack[:len(upNextStack)-1]

		i.SetRGBA(x, y, c)
		if x > xMin && colorMatches(i.RGBAAt(x-1, y), srcColor) {
			upNextStack = append(upNextStack, upNextStruct{x: x - 1, y: y})
		}
		if x < xMax && colorMatches(i.RGBAAt(x+1, y), srcColor) {
			upNextStack = append(upNextStack, upNextStruct{x: x + 1, y: y})
		}
		if y > yMin && colorMatches(i.RGBAAt(x, y-1), srcColor) {
			upNextStack = append(upNextStack, upNextStruct{x: x, y: y - 1})
		}
		if y < yMax && colorMatches(i.RGBAAt(x, y+1), srcColor) {
			upNextStack = append(upNextStack, upNextStruct{x: x, y: y + 1})
		}
	}
}

func (g *game) Update() error {
	g.controls = g.controlState.GetUserInput()

	// Pull all the items out of the command channel until it is empty.
EatDrawCommandsLoop:
	for {
		select {
		case cmd := <-g.commands:
			switch {
			case cmd.fill:
				bucketFill(g.img, cmd.x, cmd.y, cmd.c)
			case cmd.clearScreen:
				clearImage(g.img, cmd.c)
			default:
				g.img.SetRGBA(cmd.x, cmd.y, cmd.c)
			}
		default:
			// receiving from g.commands would block
			break EatDrawCommandsLoop
		}
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	// screen.DrawTriangles

	screen.WritePixels(g.img.Pix)
	if g.showFPS {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *game) getUserInput() models.UserInput {
	if g == nil || g.controls == nil {
		return models.UserInput{}
	}

	return *g.controls
}
