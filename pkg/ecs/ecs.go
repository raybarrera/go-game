package ecs

import (
	"reflect"

	"github.com/hajimehoshi/ebiten"
)

// Entity is a collection of components
type Entity struct {
	id         uint64
	components []interface{}
}

// SystemUpdater processes an update/logic on a given collection of components
type SystemUpdater interface {
	Update(screen *ebiten.Image)
}

// World manages all systems and entities
type World struct {
	SystemUpdaters []SystemUpdater
	Entities       map[reflect.Type][]reflect.Type
}

// AddSystem adds a system for the given world to manage.
func (w *World) AddSystem(system SystemUpdater) {
	w.SystemUpdaters = append(w.SystemUpdaters, system)
}

// Systems is a getter for the world's []Systems slice
func (w *World) Systems() []SystemUpdater {
	return w.Systems()
}

// Update calls update on all the systems managed by this world.
func (w *World) Update(screen *ebiten.Image) {
	for _, system := range w.SystemUpdaters {
		system.Update(screen)
	}
}

// QueryEntities returns a slice of Entities matching teh given components
func (w *World) QueryEntities(components ...reflect.Type) ([]reflect.Type, error) {
	var entities []reflect.Type
	for _, c := range components {
		for key, elem := range w.Entities {
			_, ok := containsElement(elem, c)
			if ok {
				entities = append(entities, key)
			}
		}
	}
	return entities, nil
}

func containsElement(arr []reflect.Type, check reflect.Type) (reflect.Type, bool) {
	for _, e := range arr {
		if e == check {
			return e, true
		}
	}
	return nil, false
}
