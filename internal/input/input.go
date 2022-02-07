package input

import (
	"go-game/pkg/ecs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// ActionMapComponent container for mapping actions to ebiten key presses
type ActionMapComponent struct {
	ActionMap map[ebiten.Key]func(ebiten.Key)
}

type Handler interface {
	Handle()
}

// ActionProcessorSystem is used to process all key-action maps
type ActionProcessorSystem struct {
	ecs.Entity
	KeyDownMap     ActionMapComponent
	KeyReleasedMap ActionMapComponent
}

// Update implements the interface for ecs system
func (a *ActionProcessorSystem) Update(screen *ebiten.Image) {
	for k, v := range a.KeyDownMap.ActionMap {
		if ebiten.IsKeyPressed(k) {
			v(k)
		}
	}
	for k, v := range a.KeyDownMap.ActionMap {
		if inpututil.IsKeyJustPressed(k) {
			v(k)
		}
	}
}
