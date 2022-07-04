package main

import (
	"fmt"
	"go-game/internal/game"
	"go-game/pkg/ecs"
	"go-game/pkg/tiled"
	"go-game/transform"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var TileMap tiled.TileMap
var g *game.Game
var actorSystem *game.ActorEntitySystem
var tilesImage *ebiten.Image

func init() {
	var err error
	tilesImage, _, err = ebitenutil.NewImageFromFile("Assets/Maps/gameart2d-desert.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	actors := []*game.Actor{}
	ParseMap()
	for lIndex, l := range TileMap.Layers {
		for i, t := range l.Data {
			if t == 0 {
				continue
			}
			tileSize := TileMap.TileSets[lIndex].TileWidth
			xi := float64((i % l.Width) * tileSize)
			yi := float64((i / l.Width) * tileSize)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(xi, yi)
			sub := GetSubImageFromTileset(TileMap.TileSets[lIndex], int(t))
			img := ebiten.NewImageFromImage(*sub)
			a := &game.Actor{
				Position: transform.PositionComponent{
					X: xi,
					Y: yi,
				},
				Sprite: img,
			}
			actors = append(actors, a)
		}
	}
	actorSystem = &game.ActorEntitySystem{
		Entities: actors,
	}
	g = &game.Game{
		World: &ecs.World{},
	}
	g.AddActorES(actorSystem)
}

func GetSubImageFromTileset(t tiled.TileSet, index int) *image.Image {
	xi := ((index - 1) % t.Columns) * t.TileWidth
	yi := ((index - 1) / t.Columns) * t.TileHeight
	sub := tilesImage.SubImage(image.Rect(xi, yi, xi+t.TileWidth, yi+t.TileHeight))
	return &sub
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Tiled Test Game")

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
