package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"go-game/pkg/ecs"
)

var pointerImage = ebiten.NewImage(8, 8)

// Game is a struct that contains rules and other info about a game
type Game struct {
	World *ecs.World
	Actors []*ActorEntitySystem
	Screen *ebiten.Image
}

func (g *Game) AddActorES(a *ActorEntitySystem) {
	g.Actors = append(g.Actors, a)
}

func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	//var systems []ecs.SystemUpdater
	//for _, s := range g.World.SystemUpdaters {
	//	t := reflect.TypeOf(s)
	//	if t == reflect.TypeOf(ActorEntitySystem{}) {
	//		aes := ActorEntitySystem(s)
	//	}
	//}
	for _, aes := range g.Actors {
		for _, a := range aes.Entities {
			a.Draw(screen)
		}
	}
	g.Screen = screen
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	op.GeoM.Translate(320, 240)
	screen.DrawImage(pointerImage, op)
	ebitenutil.DebugPrint(screen, "Console active")
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
