package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// MovementVelocity is used to modify movement speed
type MovementVelocity struct {
	XS float64
	YS float64
}

// MovementSystem handles movement of the given actor
type MovementSystem struct {
	Keys     []ebiten.Key
	Velocity MovementVelocity
	Actor    *Actor
}

// Update implements the ecs.SystemUpdater interface
func (ms *MovementSystem) Update(deltaTime float64) {
	for _, k := range ms.Keys {
		if ebiten.IsKeyPressed(k) {
			switch k {
			case ebiten.KeyLeft:
				ms.Actor.Position.X -= ms.Velocity.XS * deltaTime
			case ebiten.KeyRight:
				ms.Actor.Position.X += ms.Velocity.XS * deltaTime
			}
		}
	}
}
