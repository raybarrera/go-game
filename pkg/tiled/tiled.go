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
type Property struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	PropertyType string      `json:"propertytype"`
	Value        interface{} `json:"value"`
}

type TileSet struct {
	BackgroundColor  string          `json:"backgroundcolor"`
	Columns          int             `json:"columns"`
	FirstGID         int             `json:"firstgid"`
	Grid             Grid            `json:"grid"`
	Image            string          `json:"image"`
	ImageHeight      int             `json:"imageheight"`
	ImageWidth       int             `json:"imagewidth"`
	Margin           int             `json:"margin"`
	Name             string          `json:"name"`
	ObjectAlignment  string          `json:"objectalignment"`
	Properties       []Property      `json:"properties"`
	Source           string          `json:"source"`
	Spacing          int             `json:"spacing"`
	Terrains         []Terrain       `json:"terrains"`
	TileCount        int             `json:"tilecount"`
	TiledVersion     string          `json:"tiledversion"`
	TileHeight       int             `json:"tileheight"`
	TileOffset       TileOffset      `json:"tileoffset"`
	Tiles            []Tile          `json:"tiles"`
	TileWidth        int             `json:"tilewidth"`
	Transformations  Transformations `json:"transformations"`
	TransparentColor string          `json:"transparentcolor"`
	Type             string          `json:"type"`
	Version          string          `json:"version"`
}

type Tile struct {
	Animation   []Frame    `json:"animation"`
	Id          int        `json:"id"`
	Image       string     `json:"image"`
	ImageHeight int        `json:"imageheight"`
	ImageWidth  int        `json:"imagewidth"`
	ObjectGroup Layer      `json:"objectgroup"`
	Probability float32    `json:"probability"`
	Properties  []Property `json:"properties"`
	Terrain     []int      `json:"terrain"`
	Type        string     `json:"type"`
}

type Frame struct {
	Duration int `json:"duration"`
	TileId   int `json:"tileid"`
}

type Terrain struct {
	Name       string     `json:"name"`
	Properties []Property `json:"properties"`
	Tile       int        `json:"tile"`
}

type Grid struct {
	Height      int    `json:"height"`
	Width       int    `json:"width"`
	Orientation string `json:"orientation"`
}

type TileOffset struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Transformations struct {
	HFlip               bool `json:"hflip"`
	VFlip               bool `json:"vflip"`
	Rotate              bool `json:"rotate"`
	PreferUntransformed bool `json:"preferuntransformed"`
}

type Object struct {
	Ellipse    bool       `json:"ellipse"`
	GId        int        `json:"gid"`
	Height     int        `json:"height"`
	Width      int        `json:"width"`
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Point      bool       `json:"point"`
	Polygon    []Point    `json:"polygon"`
	Polyline   []Point    `json:"polyline"`
	Properties []Property `json:"properties"`
	Rotation   float64    `json:"rotation"`
	Template   string     `json:"template"`
	Text       Text       `json:"text"`
	Type       string     `json:"type"`
	Visible    bool       `json:"visible"`
	X          float64    `json:"x"`
	Y          float64    `json:"y"`
}

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Text struct{}
