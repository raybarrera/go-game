package main

import (
	"go-game/rendering"
	"go-game/transform"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var sr rendering.SpriteRenderer
var ps transform.PositionSystem

func init() {
	var img, _, _ = ebitenutil.NewImageFromFile("gopher.png", ebiten.FilterDefault)
	sr = rendering.SpriteRenderer{
		Image: img,
	}
	pos := transform.Position{
		X: 10,
		Y: 200,
	}
	ps = transform.PositionSystem{
		Sprite: sr,
		Transform: transform.Trans{
			Position: pos,
		},
	}
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	ebitenutil.DebugPrint(screen, "Hello, World")
	ps.Update(screen)
	return nil
}

func main() {
	if err := ebiten.Run(update, 640, 480, 2, "Test Game"); err != nil {
		log.Fatal(err)
	}
}
