package tiled

import (
	"encoding/json"
	"io/ioutil"
)

// TileMap implements the data structure of a tiled file as described at https://doc.mapeditor.org/en/stable/reference/json-map-format/
type TileMap struct {
	BackgroundColor  string     `json:"backgroundcolor"`
	CompressionLevel int        `json:"compressionlevel"`
	Height           int        `json:"height"`
	Width            int        `json:"width"`
	HexSideLength    int        `json:"hexsidelength"`
	Infinite         bool       `json:"infinite"`
	Layers           []Layer    `json:"layers"`
	NextLayerId      int        `json:"nextlayerid"`
	NextObjectId     int        `json:"nextobjectid"`
	Orientation      string     `json:"orientation"`
	ParallaxOriginX  float64    `json:"parallaxoriginx"`
	ParallaxOriginY  float64    `json:"parallaxoriginy"`
	Properties       []Property `json:"properties"`
	RenderOrder      string     `json:"renderorder"`
	StaggerAxis      string     `json:"staggeraxis"`
	StaggerIndex     string     `json:"staggerindex"`
	TiledVersion     string     `json:"tiledversion"`
	TileHeight       int        `json:"tileheight"`
	TileWidth        int        `json:"tilewidth"`
	TileSets         []TileSet  `json:"tilesets"`
	Type             string     `json:"type"`
	Version          string     `json:"version"`
}

// ParseMapFile takes in a path as a string, reads the file attempts to unmarshal it into a TileMap, and returns it and/or an error
func ParseMapFile(path string) (TileMap, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return TileMap{}, err
	}
	var t TileMap
	err = json.Unmarshal(bytes, &t)
	if err != nil {
		return TileMap{}, err
	}
	return t, nil
}

// ParseMapString attempts to unm
func ParseMapString(mapString string) (TileMap, error) {
	var t TileMap
	err := json.Unmarshal([]byte(mapString), &t)
	if err != nil {
		return TileMap{}, err
	}
	return t, nil
}

type Layer struct{}
type Property struct{}
type TileSet struct{}
