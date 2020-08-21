package transform

import (
	"go-game/rendering"

	"github.com/hajimehoshi/ebiten"
)

type Transform struct {
	Position Position
	Rotation Rotation
}

// Position struct containing x and y coordinates (screen-relative)
type Position struct {
	X float64
	Y float64
}

// Rotation struct containing x and y rotation values in degrees
type Rotation struct {
	X float64
	Y float64
}

// PositionSystem is the system handling placement of a sprite on screen based on its transform properties.
type PositionSystem struct {
	ImageOptions ebiten.DrawImageOptions
	Sprite       rendering.SpriteRenderer
	Transform    Transform
}

// Update method sets the rotation of the given position system and applies it to the given screen.
func (ps *PositionSystem) Update(screen *ebiten.Image) error {
	ps.ImageOptions = ebiten.DrawImageOptions{}
	ps.ImageOptions.GeoM.Translate(ps.Transform.Position.X, ps.Transform.Position.Y)
	screen.DrawImage(ps.Sprite.Image, &ps.ImageOptions)
	return nil
}
