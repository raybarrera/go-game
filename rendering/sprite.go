package rendering

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// SpriteRenderer is a entity that draws itself on screen using an image
type SpriteRenderer struct {
	TargetImage SpriteImageComponent
	Sprite      SpriteImageComponent
}

// SpriteImageComponent holds an ebiten image to bed used in a rendering sytstem
type SpriteImageComponent *ebiten.Image
