package camera

import (
	"go-game/rendering"
	"go-game/transform"
)

type RendererEntityArchetype struct {
	Transform      transform.Transform
	SpriteRenderer rendering.SpriteImageComponent
}
