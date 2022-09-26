package game

import (
	"go-game/pkg/ecs"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var pointerImage = ebiten.NewImage(8, 8)
var previousFrameTime = time.Now()

// EcsGame is a struct that contains rules and other info about a game
type EcsGame struct {
	World              *ecs.World
	ActorEntitySystems []*ActorEntitySystem
	Screen             *ebiten.Image
	ShowConsole        bool
	ScreenWidth        int
	ScreenHeight       int
	WindowWidth        int
	WindowHeight       int
}

func NewEcsGame() *EcsGame {
	return &EcsGame{
		World: &ecs.World{},
	}
}

func (g *EcsGame) AddActorES(a *ActorEntitySystem) {
	g.ActorEntitySystems = append(g.ActorEntitySystems, a)
}

func (g *EcsGame) Update() error {
	dt := float64(time.Since(previousFrameTime).Seconds())
	previousFrameTime = time.Now()
	// fmt.Printf("%v time\n", dt)
	g.World.Update(dt)
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *EcsGame) Draw(screen *ebiten.Image) {
	//var systems []ecs.SystemUpdater
	//for _, s := range g.World.SystemUpdaters {
	//	t := reflect.TypeOf(s)
	//	if t == reflect.TypeOf(ActorEntitySystem{}) {
	//		aes := ActorEntitySystem(s)
	//	}
	//}
	for _, aes := range g.ActorEntitySystems {
		for _, a := range aes.Entities {
			a.Draw(screen)
		}
	}
	g.Screen = screen
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	op.GeoM.Translate(float64(g.ScreenWidth), float64(g.ScreenHeight))
	screen.DrawImage(pointerImage, op)
	if g.ShowConsole {
		ebitenutil.DebugPrint(screen, "Console active")
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *EcsGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.WindowWidth, g.WindowHeight
}

func (g *EcsGame) toggleDebugConsole(key ebiten.Key) {
	g.ShowConsole = !g.ShowConsole
	print(g.ShowConsole)
}
