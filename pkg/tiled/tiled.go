package tiled

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

type Layer struct{}
type Property struct{}
type TileSet struct{}