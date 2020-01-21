package input

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// Mapping handles input mapping from a key to a func
type Mapping struct {
	KeysPressed map[ebiten.Key]func(ebiten.Key)
	KeysUp      map[ebiten.Key]func(ebiten.Key)
}

func GetPressedKeys() []ebiten.Key {
	var pressed []ebiten.Key
	var justReleased []ebiten.Key
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			pressed = append(pressed, k)
		}
		if inpututil.IsKeyJustPressed(k) {
			justReleased = append(justReleased, k)
		}
	}
	return pressed
}

func GetReleasedKeys() []ebiten.Key {
	var justReleased []ebiten.Key
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if inpututil.IsKeyJustPressed(k) {
			justReleased = append(justReleased, k)
		}
	}
	return justReleased
}

func (imap *Mapping) ProcessPressedKeys(keys []ebiten.Key) error {
	for _, key := range keys {
		if val, ok := imap.KeysPressed[key]; ok {
			val(key)
		}
	}
	return nil
}

func (imap *Mapping) ProcessedReleasedKeys(keys []ebiten.Key) error {
	for _, key := range keys {
		if val, ok := imap.KeysUp[key]; ok {
			val(key)
		}
	}
	return nil
}
