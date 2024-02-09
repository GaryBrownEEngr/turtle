package ebitencanvas

import (
	"bytes"
	"encoding/base64"
	"image"
	"sync"

	"github.com/GaryBrownEEngr/turtle/turtlemodel"
	"github.com/hajimehoshi/ebiten/v2"
)

type spriteToDraw struct {
	m           *sync.Mutex
	img         image.Image
	ImageEbiten *ebiten.Image
	width       int
	height      int
	x           float64
	y           float64
	angle       float64 // Radians
	visible     bool
	scale       float64
}

var _ turtlemodel.Sprite = &spriteToDraw{} // force the compiler to verify this implements the interface.

// Create a new sprite.
// Default to the turtle image and invisible.
func NewSprite() *spriteToDraw {
	ret := &spriteToDraw{
		m:       &sync.Mutex{},
		x:       0,
		y:       0,
		angle:   0,
		visible: false,
		scale:   1,
	}

	ret.SetSpriteImageTurtle()
	return ret
}

// Set the sprite bitmap image to the provided image.
func (s *spriteToDraw) SetSpriteImage(in image.Image) {
	s.m.Lock()
	defer s.m.Unlock()
	s.img = in
	s.ImageEbiten = ebiten.NewImageFromImage(in)
	bounds := in.Bounds()
	s.width = bounds.Max.X
	s.height = bounds.Max.Y
	s.scale = 1
}

// Set the sprite bitmap image to the built in turtle.
func (s *spriteToDraw) SetSpriteImageTurtle() {
	fileData, err := base64.StdEncoding.DecodeString(turtleImage)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(bytes.NewReader(fileData))
	if err != nil {
		panic(err)
	}
	s.SetSpriteImage(img)
	s.SetScale(.35)
}

// Set the sprite bitmap image to the built in arrow.
func (s *spriteToDraw) SetSpriteImageArrow() {
	fileData, err := base64.StdEncoding.DecodeString(arrowImage)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(bytes.NewReader(fileData))
	if err != nil {
		panic(err)
	}
	s.SetSpriteImage(img)
	s.SetScale(.5)
}

// Set the sprites current angle in radians.
func (s *spriteToDraw) SetRotation(radianAngle float64) {
	s.m.Lock()
	defer s.m.Unlock()
	s.angle = radianAngle
}

// Set the sprites current x/y location. 0,0 is the center of the screen.
func (s *spriteToDraw) SetPosition(cartX, cartY float64) {
	s.m.Lock()
	defer s.m.Unlock()
	s.x = cartX
	s.y = cartY
}

// Make the sprite visible on the screen.
func (s *spriteToDraw) SetVisible(visible bool) {
	s.m.Lock()
	defer s.m.Unlock()
	s.visible = visible
}

// Make the sprite scaling factor.
func (s *spriteToDraw) SetScale(scale float64) {
	s.m.Lock()
	defer s.m.Unlock()
	s.scale = scale
}

// Set all the commonly updated parameters at once. visible, x, y, angle.
func (s *spriteToDraw) Set(visible bool, cartX, cartY, radianAngle float64) {
	s.m.Lock()
	defer s.m.Unlock()
	s.visible = visible
	s.x = cartX
	s.y = cartY
	s.angle = radianAngle
}

// Get all the commonly updated parameters at once.
func (s *spriteToDraw) Get() turtlemodel.SpriteInfo {
	s.m.Lock()
	defer s.m.Unlock()
	ret := turtlemodel.SpriteInfo{
		X:       s.x,
		Y:       s.y,
		Angle:   s.angle,
		Visible: s.visible,
		Img:     s.img,
		Scale:   s.scale,
	}

	return ret
}
