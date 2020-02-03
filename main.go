package main

import (
	"go-game/internal/game"
	"go-game/internal/input"
	"go-game/rendering"
	"go-game/transform"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var imap input.Mapping
var consoleIsOpen bool

var gameScreen *ebiten.Image
var gopherActor game.ActorEntity
var actorSystem game.ActorEntitySystem

//Time info
var lastFrame float64

func main() {
	if err := ebiten.Run(update, 640, 480, 2, "Test Game"); err != nil {
		log.Fatal(err)
	}
}

// Init initialiezes the world for now
func init() {
	goa := createGopher()
	print("%v", goa.Sprite.Image)
	actorSystem = game.ActorEntitySystem{
		Entities: []game.ActorEntity{
			goa,
		},
	}
	setupInput()
}

func createGopher() game.ActorEntity {
	var img, _, err = ebitenutil.NewImageFromFile("gopher.png", ebiten.FilterDefault)
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

	gopherActor = game.ActorEntity{
		Position: pos,
		Sprite:   gopherSprite,
	}
	return gopherActor
}

func update(screen *ebiten.Image) error {
	// Time calculations
	gameScreen = screen
	now := time.Now().UnixNano()
	nowMilliseconds := now / 1000000
	dt := float64(nowMilliseconds) - lastFrame
	lastFrame = float64(time.Now().UnixNano())

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if consoleIsOpen {
		ebitenutil.DebugPrint(screen, "Hello, World")
	}
	pressedkeys := input.GetPressedKeys()
	releasedkeys := input.GetReleasedKeys()
	imap.ProcessPressedKeys(pressedkeys)
	imap.ProcessedReleasedKeys(releasedkeys)
	// ps.Update(screen)
	// renderSystem.Update(dt)
	actorSystem.Update(screen, dt)
	return nil
}

func setupInput() {
	imap = input.Mapping{
		KeysPressed: map[ebiten.Key]func(ebiten.Key){
			ebiten.KeyLeft:  handleArrows,
			ebiten.KeyRight: handleArrows,
		},
		KeysUp: map[ebiten.Key]func(ebiten.Key){
			ebiten.KeyGraveAccent: toggleDebugConsole,
		},
	}
}

func handleArrows(key ebiten.Key) {
	// switch key {
	// case ebiten.KeyLeft:
	// 	ps.Transform.Position.X--
	// case ebiten.KeyRight:
	// 	ps.Transform.Position.X++
	// }
}

func toggleDebugConsole(key ebiten.Key) {
	consoleIsOpen = !consoleIsOpen
	print(consoleIsOpen)
}
