package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"go-game/internal/game"
	"go-game/pkg/ecs"
	"go-game/pkg/tiled"
	"go-game/transform"
	"image"
	"log"
)

var TileMap tiled.TileMap
var g *game.Game
var actorSystem *game.ActorEntitySystem
var tilesImage *ebiten.Image

func init() {
	tilesImage, _, err := ebitenutil.NewImageFromFile("Assets/Maps/gameart2d-desert.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	var actors []*game.Actor
	ParseMap()
	for lIndex, l := range TileMap.Layers {
		for i, t := range l.Data {
			if t == 0 {
				continue
			}
			tileSize := TileMap.TileSets[lIndex].TileWidth
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%l.Width)*tileSize), float64((i/l.Width)*tileSize))
			sx := ((int)(t-1) % TileMap.TileSets[lIndex].Columns) * tileSize
			sy := ((int)(t-1) / TileMap.TileSets[lIndex].Columns) * tileSize
			sub := tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize))
			img := ebiten.NewImageFromImage(sub)
			a := &game.Actor{
				Position: transform.PositionComponent{},
				Sprite:   img,
			}
			actors = append(actors, a)
		}
	}
	actorSystem = &game.ActorEntitySystem{
		Entities: actors,
	}
	g.AddActorES(actorSystem)
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Tiled Test Game")
	g = &game.Game{
		World: &ecs.World{},
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func ParseMap() {
	var err error
	TileMap, err = tiled.ParseMapFile("Assets/Maps/test.tmj")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(TileMap.String())
}
