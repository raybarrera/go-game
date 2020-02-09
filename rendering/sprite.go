package rendering

import (
	"go-game/pkg/ecs"

	"github.com/hajimehoshi/ebiten"
)

// SpriteRenderer is a entity that draws itself on screen using an image
type SpriteRenderer struct {
	ecs.Entity
	TargetImage SpriteImageComponent
	Sprite      SpriteImageComponent
}

// SpriteImageComponent holds an ebiten image to bed used in a rendering sytstem
type SpriteImageComponent struct {
	Image *ebiten.Image
}
