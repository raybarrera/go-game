package main

import (
	"fmt"
	"go-game/internal/camera"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/math/f64"
)

var game *Game
var gopherImage *ebiten.Image

type Game struct {
	Cms CameraMovementSystem
}

func init() {
	var err error
	gopherImage, _, err = ebitenutil.NewImageFromFile("../../gopher.png")
	if err != nil {
		fmt.Println(err)
	}
	cms := CameraMovementSystem{
		cameraData: &camera.Camera{},
		mover: &CameraMover{
			targetPosition: f64.Vec2{
				200, 1,
			},
			speed: 0.01,
		},
	}
	cms.cameraData.Reset()
	game = &Game{
		Cms: cms,
	}
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Camera System Test")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	fps := ebiten.CurrentTPS()
	fmt.Println(fps)
	if fps < 1 {
		fps = 1
	}
	dt := 1.00 / fps
	g.Cms.Update(dt)
	fmt.Println(dt)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Cms.cameraData.Render(gopherImage, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

type CameraMover struct {
	targetPosition f64.Vec2
	speed          float64
}

type CameraMovementSystem struct {
	cameraData *camera.Camera
	mover      *CameraMover
}

func (c *CameraMovementSystem) Update(dt float64) {
	direction := f64.Vec2{
		(c.mover.targetPosition[0] - c.cameraData.Position[0]) * dt * c.mover.speed,
		(c.mover.targetPosition[1] - c.cameraData.Position[1]) * dt * c.mover.speed,
	}
	fmt.Println(direction)
	c.cameraData.Position[0] += direction[0]
	c.cameraData.Position[1] += direction[1]
}
