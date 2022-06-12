package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-game/internal/game"
	"go-game/pkg/ecs"
	"log"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Tiled Test Game")
	g := &game.Game{
		World: &ecs.World{},
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
