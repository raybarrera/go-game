package main

import (
	"go-game/internal/game"
	"go-game/pkg/ecs"
	"go-game/rendering"
	"go-game/transform"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
)

var consoleIsOpen bool

var gameScreen *ebiten.Image
var gopherActor *game.Actor
var actorSystem *game.ActorEntitySystem

var world ecs.World

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Test Game")
	g := &game.Game{}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
	//if err := ebiten.Run(update, 640, 480, 2, "Test Game"); err != nil {
	//	log.Fatal(err)
	//}
}

// Init initializes the world for now
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
	gopherSprite := rendering.SpriteImageComponent{
		Image: img,
	}
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

func update(screen *ebiten.Image) error {
	gameScreen = screen

	if consoleIsOpen {
		ebitenutil.DebugPrint(screen, "Console active")
	}
	world.Update(screen)
	return nil
}

func setupInput() *game.MovementSystem {
	return &game.MovementSystem{
		Velocity: game.MovementVelocity{
			XS: 5.0,
			YS: 5.0,
		},
		Actor: gopherActor,
		Keys: []ebiten.Key{
			ebiten.KeyLeft,
			ebiten.KeyRight,
		},
	}
}

func toggleDebugConsole(key ebiten.Key) {
	consoleIsOpen = !consoleIsOpen
	print(consoleIsOpen)
}
