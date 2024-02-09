package fakes

import (
	"image"

	"github.com/GaryBrownEEngr/turtle/turtlemodel"
)

type SpriteToDraw struct {
	Img          image.Image
	Width        int
	Height       int
	X            float64
	Y            float64
	Angle        float64
	Visible      bool
	Scale        float64
	CurrentImage string
}

var _ turtlemodel.Sprite = &SpriteToDraw{}

func NewSprite() *SpriteToDraw {
	ret := &SpriteToDraw{
		X:            0,
		Y:            0,
		Angle:        0,
		Visible:      false,
		Scale:        1,
		CurrentImage: "StartImage",
		Img:          image.NewRGBA(image.Rect(0, 0, 1, 1)),
	}

	return ret
}

func (s *SpriteToDraw) SetSpriteImage(in image.Image) {
	s.Img = in
	bounds := in.Bounds()
	s.Width = bounds.Max.X
	s.Height = bounds.Max.Y
	s.Scale = 1
	s.CurrentImage = "custom"
}

func (s *SpriteToDraw) SetSpriteImageTurtle() {
	s.CurrentImage = "turtle"
}

func (s *SpriteToDraw) SetSpriteImageArrow() {
	s.CurrentImage = "arrow"
}

func (s *SpriteToDraw) SetRotation(radianAngle float64) {
	s.Angle = radianAngle
}

func (s *SpriteToDraw) SetPosition(cartX, cartY float64) {
	s.X = cartX
	s.Y = cartY
}

func (s *SpriteToDraw) SetVisible(visible bool) {
	s.Visible = visible
}

func (s *SpriteToDraw) SetScale(scale float64) {
	s.Scale = scale
}

func (s *SpriteToDraw) Set(visible bool, cartX, cartY, radianAngle float64) {
	s.Visible = visible
	s.X = cartX
	s.Y = cartY
	s.Angle = radianAngle
}

func (s *SpriteToDraw) Get() turtlemodel.SpriteInfo {
	ret := turtlemodel.SpriteInfo{
		X:       s.X,
		Y:       s.Y,
		Angle:   s.Angle,
		Visible: s.Visible,
		Img:     s.Img,
		Scale:   s.Scale,
	}

	return ret
}
