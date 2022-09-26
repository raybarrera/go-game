package main

import (
	"fmt"
	"go-game/internal/camera"
	"go-game/internal/game"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/math/f64"
)

var gogame *game.EcsGame
var gopherImage *ebiten.Image
var previousFrameTime = time.Now()

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
	gogame = game.NewEcsGame()
	gogame.World.AddSystem(&cms)
	gogame.WindowWidth = 1280
	gogame.WindowHeight = 720

}

func main() {
	ebiten.SetWindowSize(gogame.WindowWidth, gogame.WindowHeight)
	ebiten.SetWindowTitle("ECS Camera System Test")
	if err := ebiten.RunGame(gogame); err != nil {
		log.Fatal(err)
	}
}
