package input

import "github.com/hajimehoshi/ebiten"

// Mapping handles input mapping from a key to a func
type Mapping struct {
	Keys map[ebiten.Key]func(ebiten.Key)
}

func GetPressedKeys() []ebiten.Key {
	var pressed []ebiten.Key

	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			pressed = append(pressed, k)
		}
	}
	return pressed
}

func (imap *Mapping) ProcessPressedKeys(keys []ebiten.Key) error {
	for _, key := range keys {
		if val, ok := imap.Keys[key]; ok {
			val(key)
		}
	}
	return nil
}
