package main

import (
	"math/rand"
	"time"

	"github.com/GaryBrownEEngr/turtle/ebitencanvas"
	"github.com/GaryBrownEEngr/turtle/models"
	"github.com/GaryBrownEEngr/turtle/turtleutil"
)

// Paints the screen green, then starts randomly selecting pixels and either painting them blue or black.
func main() {
	params := ebitencanvas.CanvasParams{Width: 1000, Height: 1000, ShowFPS: true}
	ebitencanvas.StartEbitenTurtleCanvas(params, drawFunc)
}

func drawFunc(can models.Canvas) {
	time.Sleep(time.Second * 3)
	can.ClearScreen(turtleutil.Green)

	width, height := can.GetWidth(), can.GetHeight()

	// st := time.Now()

	for t := 0; t < 60*60; t++ {
		for i := 0; i < 3000; i++ {
			w := rand.Intn(width)
			h := rand.Intn(height)
			whichColor := rand.Intn(2)
			c := turtleutil.Water
			if whichColor == 1 {
				c = turtleutil.Black
			}

			can.SetPixel(w, h, c)
		}
		time.Sleep(time.Second / 75)
		// et := time.Now()
		// deltaTime := et.Sub(st)
		// refreshRate := 1.0 / deltaTime.Seconds()
		// st = et
		// fmt.Printf("%5.2f\n", refreshRate)
	}

}
