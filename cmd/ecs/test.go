package main

import (
	"fmt"
	"go-game/internal/camera"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/math/f64"
)

var game *Game
var gopherImage *ebiten.Image
var previousFrameTime = time.Now()

type Game struct {
	Cms camera.CameraMovementSystem
}

func init() {
	var err error
	gopherImage, _, err = ebitenutil.NewImageFromFile("../../gopher.png")
	if err != nil {
		fmt.Println(err)
	}
	cms := camera.CameraMovementSystem{
		CameraData: &camera.Camera{},
		Mover: &camera.Mover{
			TargetPosition: f64.Vec2{
				-10, -2,
			},
			Velocity: 2,
		},
	}
	cms.CameraData.Reset()
	game = &Game{
		Cms: cms,
	}
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("ECS Camera System Test")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	dt := float64(time.Since(previousFrameTime).Seconds())
	previousFrameTime = time.Now()
	g.Cms.Update(dt)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Cms.CameraData.Render(gopherImage, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}
