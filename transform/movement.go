package transform

import (
	"go-game/rendering"

	"github.com/hajimehoshi/ebiten"
)

type PositionSystem struct {
	ImageOptions ebiten.DrawImageOptions
	Sprite       rendering.SpriteRenderer
	Transform    Trans
}

func (ps *PositionSystem) Update(screen *ebiten.Image) error {
	ps.ImageOptions = ebiten.DrawImageOptions{}
	ps.ImageOptions.GeoM.Translate(ps.Transform.Position.X, ps.Transform.Position.Y)
	screen.DrawImage(ps.Sprite.Image, &ps.ImageOptions)
	return nil
}
