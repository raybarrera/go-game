package camera

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
)

type Camera struct {
	ViewPort   f64.Vec2
	Position   f64.Vec2
	ZoomFactor int
	Rotation   int
}

func (c *Camera) String() string {
	return fmt.Sprintf("T: %.1f, R: %d, S: %d", c.Position, c.Rotation, c.ZoomFactor)
}

func (c *Camera) ViewPortCenter() f64.Vec2 {
	return f64.vec2{
		c.ViewPort[0] * 0.5,
		c.ViewPort[1] * 0.5,
	}
}

func (c *Camera) WorldMatrix() ebiten.GeoM {
	m := ebiten.GeoM{}
	m.Translate(-c.Position[0], -c.Position[1])
	m.Translate(-c.ViewPortCenter()[0], -c.ViewPortCenter()[1])
	m.Scale(
		math.Pow(1.01, float64(c.ZoomFactor)),
		math.Pow(1.01, float64(c.ZoomFactor)),
	)
	//Rotate
	m.Translate(c.ViewPortCenter()[0], c.ViewPortCenter()[1])
	return m
}

func (c *Camera) Render(world, screen *ebiten.Image) {
	screen.DrawImage(world, &ebiten.DrawImageOptions{
		GeoM: c.WorldMatrix(),
	})
}

func (c *Camera) ScreenToWorld(posX, posY int) (float64, float64) {
	inverseMatrix := c.WorldMatrix()
	if inverseMatrix.IsInvertible() {
		inverseMatrix.Invert()
		return inverseMatrix.Apply(float64(posX), float64(posY))
	} else {
		return math.NaN(), math.NaN()
	}
}

func (c *Camera) Reset() {
	c.Position[0] = 0
	c.Position[1] = 0
	c.Rotation = 0
	c.ZoomFactor = 0
}
