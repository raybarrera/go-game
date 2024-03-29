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
	return f64.Vec2{
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

type Mover struct {
	TargetPosition f64.Vec2
	Velocity       float64
}

type CameraMovementSystem struct {
	//TODO This code needs to be moved so that the Update method is in charge of fetching all the matching entities with matching types. Right now, we'd need to create a system instance for each camera (or entity) we want to update. Instead, the system should find and operate on the entities on its own.
	CameraData *Camera
	Mover      *Mover
}

func (c *CameraMovementSystem) Update(dt float64) {
	c.CameraData.Position[0] += c.Mover.TargetPosition[0] * dt * c.Mover.Velocity
	c.CameraData.Position[1] += c.Mover.TargetPosition[1] * dt * c.Mover.Velocity
}
