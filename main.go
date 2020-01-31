package main

import (
	"go-game/internal/input"
	"go-game/rendering"
	"go-game/transform"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var sr rendering.SpriteRenderer
var ps transform.PositionSystem
var imap input.Mapping
var consoleIsOpen bool

var gopher rendering.SpriteRenderer
var renderSystem rendering.SpriteRenderSystem

var gameScreen *ebiten.Image

//Time info
var lastFrame float64

func main() {
	if err := ebiten.Run(update, 640, 480, 2, "Test Game"); err != nil {
		log.Fatal(err)
	}
	Init()
}

func Init() {
	createGopher()
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

func createGopher() {
	var img, _, _ = ebitenutil.NewImageFromFile("gopher.png", ebiten.FilterDefault)
	gopherSprite := rendering.SpriteImageComponent{
		Image: img,
	}
	targetSprite := rendering.SpriteImageComponent{
		Image: gameScreen,
	}
	gopher = rendering.SpriteRenderer{
		TargetImage: targetSprite,
		Sprite:      gopherSprite,
	}
	renderSystem = rendering.SpriteRenderSystem{
		Entities: []rendering.SpriteRenderer{
			gopher,
		},
	}
}

func update(screen *ebiten.Image) error {
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
	ps.Update(screen)
	renderSystem.Update(dt)

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
	switch key {
	case ebiten.KeyLeft:
		ps.Transform.Position.X--
	case ebiten.KeyRight:
		ps.Transform.Position.X++
	}
}

func toggleDebugConsole(key ebiten.Key) {
	consoleIsOpen = !consoleIsOpen
	print(consoleIsOpen)
}
