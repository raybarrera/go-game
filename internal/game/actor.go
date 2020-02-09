package game

import (
	"go-game/pkg/ecs"
	"go-game/rendering"
	"go-game/transform"

	"github.com/hajimehoshi/ebiten"
)

// Actor represents a game object in the world.
type Actor struct {
	ecs.Entity
	Position    transform.PositionComponent
	Rotation    transform.RotationComponent
	TargetImage rendering.SpriteImageComponent
	Sprite      rendering.SpriteImageComponent
}

// ActorEntitySystem draws all actors at a given position
type ActorEntitySystem struct {
	Entities []*Actor
}

//Update draws one frame of the actor
func (e *ActorEntitySystem) Update(screen *ebiten.Image) {
	if ebiten.IsDrawingSkipped() {
		return
	}
	for _, entity := range e.Entities {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(entity.Position.X, entity.Position.Y)
		screen.DrawImage(entity.Sprite.Image, options)
	}
}
