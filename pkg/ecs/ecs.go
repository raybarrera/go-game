package ecs

import (
	"github.com/google/uuid"
	"reflect"
)

// Entity is a collection of components
// TODO possibly need to remove this, or adopt it more generally. Might constrain entities too much to fit this container.
// the alternative is to rely on reflection to get entities, which could be any type without this constraint.
// There is a possibility of using an interface here as well, but it feels a bit forced. - Ray.
type Entity uuid.UUID

// NewEntity returns an instance of Entity with a valid internal UUID
func NewEntity() *Entity {
	//Facade the usage of the uuid package, which is itself a byte slice
	id, err := uuid.NewUUID()
	if err != nil {
		return nil
	}
	e := Entity(id)
	return &e
}

// SystemUpdater processes an update/logic on a given collection of components
type SystemUpdater interface {
	Update()
}

// World manages all systems and entities
// TODO: Entities is not in use. Right now entities are arrays inside of systems, not the world. Pick a lane.
type World struct {
	SystemUpdaters []SystemUpdater
	Entities       map[interface{}][]reflect.Type
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
func (w *World) Update() {
	for _, system := range w.SystemUpdaters {
		system.Update()
	}
}

// QueryEntities returns a slice of Entities matching the given components
//
// This functionality is loosely based on Unity's ECS EntityQuery implementation
// albeit, purely based on the public API since AFAIK, the implementation is closed-source.
func (w *World) QueryEntities(components ...reflect.Type) ([]interface{}, error) {
	var entities []interface{}
	for _, c := range components {
		for key, elem := range w.Entities {
			_, ok := ContainsElement(elem, c)
			if ok {
				entities = append(entities, key)
			}
		}
	}
	return entities, nil
}

type EntityManager struct {
	// Entities maps an entity (an uuid, essentially) to a slice of components (data/structs)
	Entities map[Entity][]interface{}
}

// ContainsElement is a helper function that finds the given type in the type array.
func ContainsElement(arr []reflect.Type, check reflect.Type) (reflect.Type, bool) {
	for _, e := range arr {
		if e == check {
			return e, true
		}
	}
	return nil, false
}
