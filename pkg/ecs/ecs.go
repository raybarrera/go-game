package ecs

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

type Entity uuid.UUID

func (id Entity) String() string {
	return uuid.UUID(id).String()
}

func NewEntity() Entity {
	id, _ := uuid.NewUUID()
	return Entity(id)
}

type EntityCollection struct {
}

// SystemUpdater processes an update/logic on a given collection of components
type SystemUpdater interface {
	Update(deltaTime float64)
}

// World manages all systems and entities
// TODO: Entities is not in use. Right now entities are arrays inside of systems, not the world. Pick a lane.
type World struct {
	SystemUpdaters []SystemUpdater
	EntityManager  EntityManager
}

func NewWorld() *World {
	return &World{
		SystemUpdaters: make([]SystemUpdater, 0, 10),
		EntityManager: EntityManager{
			Entities: map[Entity][]interface{}{},
		},
	}
}

// AddSystem adds a system for the given world to manage.
func (w *World) AddSystem(system SystemUpdater) {
	w.SystemUpdaters = append(w.SystemUpdaters, system)
}

// Systems is a getter for the world's []Systems slice
func (w *World) Systems() []SystemUpdater {
	return w.Systems()
}

func (w *World) CreateEntity(components []interface{}) {
	w.EntityManager.Entities[NewEntity()] = components
}

// Update calls update on all the systems managed by this world.
func (w *World) Update(deltaTime float64) {
	for _, system := range w.SystemUpdaters {
		system.Update(deltaTime)
	}
}

func (w *World) AddComponent(component interface{}, toEntity Entity) {}

func (w *World) modifyRegisteredEntity(entity Entity, newTemplate ...reflect.Type) {}

// QueryEntities returns a slice of Entities matching the given components
//
// This functionality is loosely based on Unity's ECS EntityQuery implementation
// albeit, purely based on the public API since AFAIK, the implementation is closed-source.
func (w *World) QueryEntities(components ...reflect.Type) (EntityCollection, error) {
	var matchingEntities []Entity
	var ok bool
	for key, e := range w.EntityManager.Entities {
		for _, c := range components {
			ok = ContainsType(e, c)
			if !ok {
				break
			}
		}
		matchingEntities = append(matchingEntities, key)
	}
	return EntityCollection{}, nil

}

// Template contains a combination of types shared by various Entities.
// Definition maps a type to a slice of elements of that type.
// The definition keys array can be used to query based on component types.
type Template struct {
	Id         int
	Definition map[reflect.Type][]interface{}
}

// componentStore maps the addresses of its components in the template store.
type componentStore struct {
	Entities map[Entity][]int
}

type EntityManager struct {
	// Entities maps an entity (an uuid, essentially) to a slice of components (data/structs)
	Entities map[Entity][]interface{}
}

func hash(components ...interface{}) []byte {
	var b bytes.Buffer
	err := gob.NewEncoder(&b).Encode(components)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return b.Bytes()
}

func ContainsType(arr []interface{}, check reflect.Type) bool {
	for _, e := range arr {
		if reflect.TypeOf(e) == check {
			return true
		}
	}
	return false
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
