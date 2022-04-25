package game

import (
	"go-game/rendering"
	"go-game/transform"

	"github.com/hajimehoshi/ebiten/v2"
)

// Actor represents a game object in the world.
type Actor struct {
	Position    transform.PositionComponent
	Rotation    transform.RotationComponent
	TargetImage rendering.SpriteImageComponent
	Sprite      rendering.SpriteImageComponent
}

// NewActor Returns an *Actor with a valid UUID
func NewActor() *Actor {
	return &Actor{}
}

func (a Actor) Draw(image *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(a.Position.X, a.Position.Y)
	image.DrawImage(a.Sprite, options)
}

// ActorEntitySystem draws all actors at a given position
type ActorEntitySystem struct {
	Entities []*Actor
}

//Update draws one frame of the actor
func (e *ActorEntitySystem) Update(deltaTime float64) {
	//for _, entity := range e.Entities {
	//	entity.Draw(g.Screen)
	//}
}
