package main

import (
	"image/color"
	"time"

	"github.com/GaryBrownEEngr/turtle"
)

// Draws The Go Gopher.
// Converted from the python script seen in this youtube video: https://www.youtube.com/watch?v=d8A1jqOGzNE
// Since this turtle system doesn't have poly-fill like python does, bucket fill is used instead.
func main() {
	params := turtle.Params{Width: 500, Height: 600}
	turtle.Start(params, drawFunc)
}

func drawFunc(window turtle.Window) {
	blue := color.RGBA{0x74, 0xCE, 0xDD, 0xFF}  // #74CEDD
	brown := color.RGBA{0xF7, 0xD3, 0xA2, 0xFF} // #F7D3A2
	black := color.RGBA{0x00, 0x00, 0x00, 0xFF}
	black2 := color.RGBA{1, 1, 1, 0xFF}
	White := color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}

	t := window.NewTurtle()
	t.ShapeAsArrow()
	t.ShowTurtle()
	t.Speed(100)
	t.Size(4)

	// go turtle.CreateGifMostCommonColors(window, time.Millisecond*600, time.Millisecond*100, "./examples/gogopher/GoGopher.gif", int(60/.6))

	t.R(57.8)
	t.F(181.3)
	t.On()
	t.Color(black)

	// Right foot
	t.L(14.7)
	t.F(13.5)
	t.R(5.4)
	t.F(12.6)
	t.R(14.5)
	t.F(7.6)
	t.R(28.4)
	t.F(4.8)
	t.R(32.2)
	t.F(6.6)
	t.R(21.1)
	t.F(7.7)
	t.R(36.5)
	t.F(5.4)
	t.R(33)
	t.F(10.3)
	t.R(16.6)
	t.F(9.3)
	t.L(9.3)
	t.F(8.7)
	t.R(1.8)
	t.F(9.1)
	t.R(101.6)
	t.F(26.1)
	t.Off()

	t.R(104.2)
	t.F(25.9)
	t.Fill(brown)
	t.On()
	t.L(14.7)
	t.F(6.3)
	t.R(14.4)
	t.F(8.2)
	t.Off()

	// Left foot
	t.R(115.4)
	t.F(181.5)
	t.On()
	t.L(11.4)
	t.F(8.8)
	t.L(19.7)
	t.F(10.1)
	t.L(11.6)
	t.F(11.5)
	t.L(2)
	t.F(6.6)
	t.R(15.4)
	t.F(6.4)
	t.R(20.8)
	t.F(4.9)
	t.R(40.9)
	t.F(4.7)
	t.R(26)
	t.F(3.5)
	t.L(58.7)
	t.F(1.8)
	t.R(49.9)
	t.F(4.6)
	t.R(30.2)
	t.F(4.7)
	t.R(19.8)
	t.F(6.2)
	t.R(28.6)
	t.F(11.3)
	t.R(10)
	t.F(17.7)
	t.R(70.8)
	t.F(35.5)
	t.Off()
	t.R(144.1)
	t.F(40.6)
	t.Fill(brown)

	t.On()
	t.L(58.1)
	t.F(6)
	t.R(9.1)
	t.F(5.2)
	t.L(17.7)
	t.F(3.1)
	t.Off()

	// Right arm
	t.L(151.7)
	t.F(322.1)
	t.On()
	t.R(42.7)
	t.F(11)
	t.R(17.6)
	t.F(9.2)
	t.R(14.7)
	t.F(6.7)
	t.R(17.4)
	t.F(7.2)
	t.R(26.6)
	t.F(4.4)
	t.R(60.1)
	t.F(4.1)
	t.L(44.8)
	t.F(3.7)
	t.R(34.3)
	t.F(3.8)
	t.R(36.4)
	t.F(4.8)
	t.R(29.2)
	t.F(5.8)
	t.R(16.5)
	t.F(7.8)
	t.L(9.9)
	t.F(9.1)
	t.R(71.8)
	t.F(19.5)
	t.Off()

	t.R(125.6)
	t.F(27.8)
	t.Fill(brown)
	t.On()
	t.R(5.7)
	t.F(3.7)
	t.L(34.4)
	t.F(2.8)
	t.Off()

	// Left arm
	t.R(175.2)
	t.F(277.6)
	t.On()
	t.R(1.4)
	t.F(9.7)
	t.L(16.7)
	t.F(9.8)
	t.L(18.3)
	t.F(10.3)
	t.L(16.7)
	t.F(6.1)
	t.L(28.2)
	t.F(4.7)
	t.L(49.2)
	t.F(5)
	t.R(21.7)
	t.F(4)
	t.L(27.1)
	t.F(4.5)
	t.L(47.7)
	t.F(4.8)
	t.L(19.9)
	t.F(5.9)
	t.L(16.1)
	t.F(9.2)
	t.R(20.1)
	t.F(8.2)
	t.L(74)
	t.F(18.5)

	t.Off()
	t.L(125.4)
	t.F(30.6)
	t.Fill(brown)
	t.On()
	t.L(10.8)
	t.F(2.6)
	t.R(21.5)
	t.F(3.3)

	// right ear
	t.Off()
	t.R(161.8)
	t.F(315.1)
	t.On()
	t.L(14.1)
	t.Circle(-27.8, 93.7, 20)
	t.L(4.3)
	t.Circle(-30.6, 54.2, 20)
	t.R(3.5)
	t.Circle(-27, 72.7, 20)
	t.R(67.8)
	t.F(55.2)
	t.Off()
	t.R(153.7)
	t.F(16.7)
	t.Fill(blue)

	t.On()
	t.L(58.7)
	t.Circle(-10, 175.2, 20)
	t.Off()
	t.R(92.4)
	t.F(20)
	t.On()
	t.L(180)
	t.F(20)
	t.Off()
	t.B(10)
	t.PanL(5)
	t.Fill(black)
	t.PanR(5)
	t.F(10)

	// Left ear

	t.R(129.6)
	t.F(210.9)
	t.On()
	t.R(37.4)
	t.Circle(30.3, 86, 20)
	t.R(3.9)
	t.Circle(27.8, 56.1, 20)
	t.R(0.2)
	t.Circle(25.3, 79, 20)
	t.L(64.4)
	t.F(55.3)

	t.Off()
	t.L(152.2)
	t.F(23.1)
	t.Fill(blue)
	t.On()
	t.R(64.8)
	t.Circle(9.4, 180.7, 20)
	t.R(3.4)
	t.F(3.8)
	t.L(27.8)
	t.F(5.7)
	t.L(67.6)
	t.F(15.3)
	t.L(70.2)
	t.F(5.8)
	t.L(22.4)
	t.F(2.9)
	t.Off()
	t.PanL(10)
	t.Fill(black)
	t.PanR(10)

	// Body
	t.R(129.8)
	t.F(124.2)
	t.Color(black2) // use a slightly different Black so we can use fill to get rid of parts inside
	t.On()
	t.R(20.7)
	t.Circle(-200.6, 19.5, 20)
	t.R(3.4)
	t.Circle(-90.3, 55, 20)
	t.L(0.5)
	t.Circle(-324.4, 13.6, 20)
	t.L(0.5)
	t.Circle(2316.4, 3.4, 20)
	t.L(0.8)
	t.Circle(-136.8, 43.7, 20)
	t.R(10)
	t.Circle(-160.4, 80.8, 20)
	t.R(7.9)
	t.Circle(-105.7, 51.9, 20)
	t.L(2.6)
	t.Circle(546.9, 16, 20)
	t.R(1.9)
	t.Circle(-168.1, 36.9, 20)
	t.R(17.1)
	t.Circle(-155.8, 44.3, 20)

	t.Off()
	t.R(89.2)
	t.F(120.6)
	t.Fill(black)
	t.Fill(brown)
	t.Fill(blue)

	// Teeth
	t.Color(black)
	t.On()
	t.R(2.9)
	t.F(15.4)
	t.L(3.2)
	t.F(9.7)
	t.R(48.3)
	t.F(8.8)
	t.R(36.4)
	t.F(4.4)
	t.R(41.3)
	t.F(6.2)
	t.R(32.9)
	t.F(7.9)
	t.R(18.6)
	t.F(8.3)
	t.R(10.9)
	t.F(9.4)
	t.R(73.4)
	t.F(18.9)
	t.Off()
	t.B(9)
	t.PanR(9)
	t.Fill(White)
	t.F(9)
	t.PanL(9)

	t.R(14.5)
	t.F(16.1)
	t.On()
	t.R(82.3)
	t.F(13.8)
	t.R(0.8)
	t.F(10.9)
	t.R(47.5)
	t.F(5.9)
	t.R(22.9)
	t.F(4.4)
	t.R(42)
	t.F(5.1)
	t.R(26)
	t.F(5.6)
	t.R(34.6)
	t.F(11.3)
	t.R(11.8)
	t.F(12.4)
	t.R(71.3)
	t.F(11)
	t.R(32.5)
	t.F(6.7)
	t.Off()
	t.B(10)
	t.PanR(15)
	t.Fill(White)
	t.F(10)
	t.PanL(15)

	// Nose

	t.L(163.1)
	t.F(35.3)
	t.Color(black2)
	t.On()
	t.L(36.7)
	t.F(7.7)
	t.L(37.3)
	t.F(8.3)
	t.L(4.8)
	t.Circle(13, 96.3, 20)
	t.L(18)
	t.F(4.4)
	t.L(31.3)
	t.F(8.2)
	t.L(12.2)
	t.F(4.9)
	t.L(3.4)
	t.Circle(-23.9, 50.8, 20)
	t.R(13.6)
	t.Circle(14.3, 78.4, 20)
	t.L(1.3)
	t.Circle(10.5, 94.1, 20)
	t.R(4)
	t.Circle(32.3, 34.7, 20)
	t.L(19.5)
	t.F(26.4)
	t.Off()
	t.PanL(10)
	t.Fill(black)
	t.Fill(White)
	t.Fill(brown)
	t.PanR(10)

	// Black part of nose

	t.R(107.1)
	t.F(12.9)
	t.Color(black)
	t.On()
	t.R(49.4)
	t.Circle(-21.7, 62.5, 20)
	t.L(4.8)
	t.Circle(-8.7, 84.3, 20)
	t.R(25.3)
	t.Circle(-28.4, 37.6, 20)
	t.L(0.5)
	t.Circle(-23.2, 40.7, 20)
	t.R(17.4)
	t.Circle(-8.2, 74.8, 20)
	t.L(5)
	t.Circle(-12, 36.1, 20)
	t.Off()
	t.PanR(10)
	t.Fill(black2)
	t.Fill(brown)
	t.Fill(black)
	t.PanL(10)

	// Whites of eyes

	t.L(41.9)
	t.F(66.5)
	t.On()
	t.R(25)
	t.Circle(-42.1, 77.8, 20)
	t.R(2.4)
	t.Circle(-37.3, 81, 20)
	t.R(6.1)
	t.Circle(-45.7, 57.6, 20)
	t.L(3.1)
	t.Circle(-36.5, 73.1, 20)
	t.R(2.2)
	t.Circle(-37.2, 65.4, 20)
	t.Off()
	t.PanR(10)
	t.Fill(White)
	t.PanL(10)

	t.L(150)
	t.F(51.2)
	t.On()
	t.R(37)
	t.Circle(44.1, 67.7, 20)
	t.L(0.4)
	t.Circle(37.3, 72.1, 20)
	t.L(0.9)
	t.Circle(36.3, 78.4, 20)
	t.R(7.4)
	t.Circle(40, 66.7, 20)
	t.L(3.3)
	t.Circle(37.8, 73.8, 20)
	t.Off()
	t.PanL(10)
	t.Fill(White)
	t.PanR(10)

	// pupils
	t.R(165.6)
	t.F(50.9)
	t.On()
	t.L(71.8)
	t.Circle(-10.6, 99.1, 20)
	t.L(3.9)
	t.Circle(-12.8, 90.1, 20)
	t.L(1.5)
	t.Circle(-10.5, 85.9, 20)
	t.R(0.8)
	t.Circle(-13.6, 88.3, 20)
	t.Off()
	t.PanR(10)
	t.Fill(black)
	t.PanL(10)

	t.L(131.9)
	t.F(83.8)
	t.On()
	t.R(47.1)
	t.Circle(11.2, 91.4, 20)
	t.L(2.6)
	t.Circle(13.8, 88.5, 20)
	t.L(3.2)
	t.Circle(11.3, 76.2, 20)
	t.L(3.5)
	t.Circle(13.6, 93.4, 20)
	t.Off()
	t.PanL(10)
	t.Fill(black)
	t.PanR(10)

	// pupil reflections
	t.R(139.8)
	t.F(100.6)
	t.On()
	t.Size(2)
	t.Color(White)
	t.L(60.3)
	t.Circle(3.3, 360, 40)
	t.Off()
	t.PanL(3)
	t.Fill(White)
	t.PanR(3)

	t.L(127.1)
	t.F(99.3)
	t.On()
	t.R(95.1)
	t.Circle(3.4, 360, 40)
	t.Off()
	t.PanL(3)
	t.Fill(White)
	t.PanR(3)

	time.Sleep(time.Second)
	t.HideTurtle()

	// turtle.TakeScreenshot(window, "./examples/gogopher/GoGopher.png")
}
