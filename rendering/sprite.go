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

// SpriteRenderSystem draws an image to an image
type SpriteRenderSystem struct {
	Entities []SpriteRenderer
}

// Update draws the sprite renderer each frame
func (r *SpriteRenderSystem) Update(dt float64) {
	if ebiten.IsDrawingSkipped() {
		return
	}
	for _, entity := range r.Entities {
		entity.TargetImage.Image.DrawImage(entity.Sprite.Image, nil)
	}
}
