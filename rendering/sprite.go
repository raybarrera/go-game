package rendering

import (
	"go-game/pkg/ecs"

	"github.com/hajimehoshi/ebiten"
)

type SpriteRenderer struct {
	ecs.Component
	Image *ebiten.Image
}

func (s *SpriteRenderer) Update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	screen.DrawImage(s.Image, nil)
	return nil
}
