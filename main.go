package main

import (
	"go-game/internal/input"
	"go-game/rendering"
	"go-game/transform"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var sr rendering.SpriteRenderer
var ps transform.PositionSystem
var imap input.Mapping

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
		Transform: transform.Transform{
			Position: pos,
		},
	}
	setupInput()
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	ebitenutil.DebugPrint(screen, "Hello, World")
	pressedkeys := input.GetPressedKeys()
	imap.ProcessPressedKeys(pressedkeys)
	ps.Update(screen)
	return nil
}

func main() {
	if err := ebiten.Run(update, 640, 480, 2, "Test Game"); err != nil {
		log.Fatal(err)
	}
}

func setupInput() {
	imap = input.Mapping{
		Keys: map[ebiten.Key]func(ebiten.Key){
			ebiten.KeyLeft:  handleArrows,
			ebiten.KeyRight: handleArrows,
		},
	}
}

func handleArrows(key ebiten.Key) {
	switch key {
	case ebiten.KeyLeft:
		ps.Transform.Position.X--
	case ebiten.KeyRight:
		ps.Transform.Position.X++
	}
}
