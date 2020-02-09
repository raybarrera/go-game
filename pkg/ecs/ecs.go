package ecs

import "github.com/hajimehoshi/ebiten"

// Entity is a collection of components
type Entity struct {
	id         uint64
	components []Component
}

// Component is a data container
type Component struct {
}

// System processes an update/logic on a given collection of components
type System interface {
	Update(screen *ebiten.Image)
}

// World manages all systems and entities
type World struct {
	systems []System
}

// AddComponent adds a component to an entity
func (e *Entity) AddComponent(c Component) error {
	e.components = append(e.components, c)
	return nil
}

// AddSystem adds a system for the given world to manage.
func (w *World) AddSystem(system System) {
	w.systems = append(w.systems, system)
}

// Systems is a getter for the world's []Systems slice
func (w *World) Systems() []System {
	return w.Systems()
}

// Update calls update on all the systems managed by this world.
func (w *World) Update(screen *ebiten.Image) {
	for _, system := range w.systems {
		system.Update(screen)
	}
}
