package main

import (
	"go-game/internal/game"
	"go-game/pkg/ecs"
	"go-game/transform"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var gameScreen *ebiten.Image
var gopherActor *game.Actor
var actorSystem *game.ActorEntitySystem

var world ecs.World

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Test Game")
	g := &game.Game{
		World: &world,
	}
	g.AddActorES(actorSystem)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

// Init initializes the world and gopher for now
func init() {
	goa := createGopher()
	actorSystem = &game.ActorEntitySystem{
		Entities: []*game.Actor{
			goa,
		},
	}
	movementSystem := setupInput()
	world = ecs.World{}

	world.AddSystem(actorSystem)
	world.AddSystem(movementSystem)
}

func createGopher() *game.Actor {
	var img, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	gopherSprite := img
	pos := transform.PositionComponent{
		X: 50,
		Y: 50,
	}

	gopherActor = &game.Actor{
		Position: pos,
		Sprite:   gopherSprite,
	}
	return gopherActor
}

func setupInput() *game.MovementSystem {
	return &game.MovementSystem{
		Velocity: game.MovementVelocity{
			XS: 50.0,
			YS: 50.0,
		},
		Actor: gopherActor,
		Keys: []ebiten.Key{
			ebiten.KeyLeft,
			ebiten.KeyRight,
		},
	}
}
