package rendering

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer interface {
	Draw(*ebiten.Image)
}
