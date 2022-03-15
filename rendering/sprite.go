package rendering

import (
	"go-game/pkg/ecs"

	"github.com/hajimehoshi/ebiten/v2"
)

// SpriteRenderer is a entity that draws itself on screen using an image
type SpriteRenderer struct {
	ecs.Entity
	TargetImage SpriteImageComponent
	Sprite      SpriteImageComponent
}

// SpriteImageComponent holds an Ebiten image to be used in a rendering system
type SpriteImageComponent struct {
	Image *ebiten.Image
}

