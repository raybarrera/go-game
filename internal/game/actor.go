package game

import (
	"go-game/pkg/ecs"
	"go-game/rendering"
	"go-game/transform"

	"github.com/hajimehoshi/ebiten/v2"
)

// Actor represents a game object in the world.
type Actor struct {
	ecs.Entity
	Position    transform.PositionComponent
	Rotation    transform.RotationComponent
	TargetImage rendering.SpriteImageComponent
	Sprite      rendering.SpriteImageComponent
}

// NewActor Returns an *Actor with a valid UUID
func NewActor() *Actor {
	return &Actor{
		Entity: ecs.Entity{
			Id: 0, //Not really valid id
		},
	}
}

func (a Actor) Draw(image *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(a.Position.X, a.Position.Y)
	image.DrawImage(a.Sprite.Image, options)
}

// ActorEntitySystem draws all actors at a given position
type ActorEntitySystem struct {
	Entities []*Actor
}

//Update draws one frame of the actor
func (e *ActorEntitySystem) Update() {
	//for _, entity := range e.Entities {
	//	entity.Draw(g.Screen)
	//}
}
