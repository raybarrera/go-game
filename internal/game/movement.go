package game

import (
	"go-game/pkg/ecs"

	"github.com/hajimehoshi/ebiten"
)

// MovementVelocity is used to modify movement speed
type MovementVelocity struct {
	XS float64
	YS float64
}

// MovementSystem handles movement of the given actor
type MovementSystem struct {
	ecs.Entity
	Keys     []ebiten.Key
	Velocity MovementVelocity
	Actor    *Actor
}

// Update implements the ecs system interface
func (ms *MovementSystem) Update(screen *ebiten.Image) {
	for _, k := range ms.Keys {
		if ebiten.IsKeyPressed(k) {
			switch k {
			case ebiten.KeyLeft:
				ms.Actor.Position.X -= ms.Velocity.XS
			case ebiten.KeyRight:
				ms.Actor.Position.X += ms.Velocity.XS
			}
		}
	}
}
