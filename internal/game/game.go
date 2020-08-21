package game

import "go-game/pkg/ecs"

// Game is a struc that contains rules and other info about a game
type Game struct {
}

// World is a container for physics, entities, etc.
type World struct {
	Entities []*ecs.Entity
}

// AddEntity adds the given entity to the world's entity slice
func (w *World) AddEntity(e ecs.Entity) {
	w.Entities = append(w.Entities, &e)
	e.Start()
}
