package ecs

import (
	"fmt"
	"hash/fnv"
	"reflect"

	"github.com/google/uuid"
)

// Entity is a uuid. Conceptually, however, an entity is defined by the components it's comprised of.
type Entity uuid.UUID

func (id Entity) String() string {
	return uuid.UUID(id).String()
}

func NewEntityId() Entity {
	id, _ := uuid.NewUUID()
	return Entity(id)
}

// SystemUpdater processes an update/logic on a given collection of components
type SystemUpdater interface {
	Update(deltaTime float64)
}

// World manages all systems and entities
// TODO: Entities is not in use. Right now entities are arrays inside of systems, not the world. Pick a lane.
type World struct {
	SystemUpdaters []SystemUpdater
	ArchetypeStore map[uint32]Archetype
}

func NewWorld() *World {
	return &World{
		SystemUpdaters: make([]SystemUpdater, 0, 10),
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

func (w *World) CreateEntity(components []any) {
	h := componentsToHash(components...)
	var arch *Archetype
	if val, ok := w.ArchetypeStore[h]; ok {
		arch = &val
	} else {
		var types []reflect.Type
		for _, v := range components {
			types = append(types, reflect.TypeOf(v))
		}
		arch = NewArchetype(h, types...)
	}
	fmt.Println(arch.String())
	// store components
}

func AddComponent[T any](to Entity, component T) {
	//TODO: Calculate old archetype
	//TODO: Calculate new archetype
	//TODO: Migrate to new archetype
}

// RemoveComponentOfType TODO needs to figure out what happens when we have multiple components of the same type
func RemoveComponentOfType[T reflect.Type](from Entity, component T) {

	//TODO: Calculate old archetype
	//TODO: Calculate new archetype
	//TODO: Remove target component
	//TODO: Migrate to new archetype
}

func BatchRemoveComponent(from Entity, components ...any) {}

// Update calls update on all the systems managed by this world.
func (w *World) Update(deltaTime float64) {
	for _, system := range w.SystemUpdaters {
		system.Update(deltaTime)
	}
}

func ForEach[T any](f func(T)) {

}

// Archetype contains a combination of types shared by various Entities.
// Definition maps a type to a slice of elements of that type.
// The definition keys array can be used to query based on component types.
type Archetype struct {
	Id         uint32
	NextIndex  int
	definition map[reflect.Type][]any
}

func NewArchetype(id uint32, componentTypes ...reflect.Type) *Archetype {
	a := Archetype{
		Id:         id,
		NextIndex:  0,
		definition: make(map[reflect.Type][]any),
	}
	for _, v := range componentTypes {
		a.definition[v] = make([]any, 0)
	}
	return &a
}

func (a *Archetype) String() string {
	s := ""
	s += fmt.Sprintf("Archetype ID: %v\n", a.Id)
	for k, v := range a.definition {
		s += fmt.Sprintf("| type: %v | items: %v |\n", k.String(), len(v))
	}
	s += fmt.Sprintf("| NextIndex: %v | Valid: %v |\n", a.NextIndex, a.GetNextIndex())
	return s
}

func (a *Archetype) AddEntity(components []any) {
	appendMode := false
	if isFull := a.GetNextIndex() == -1; isFull {
		appendMode = true
	}
	for _, v := range components {
		t := reflect.TypeOf(v)
		if appendMode {
			a.definition[t] = append(a.definition[t], v)
		} else {
			a.definition[t][a.NextIndex] = v
		}

	}
	a.NextIndex++
}

func (a *Archetype) GetNextIndex() int {
	next := a.NextIndex
	for _, v := range a.definition {
		if len(v) <= 0 {
			break
		}
		if v[next] == nil {
			break
		}
		isFull := true
		for i, elem := range v {
			if elem == nil {
				isFull = false
				next = i
				break
			}
		}
		if isFull {
			a.NextIndex = len(v)
			next = -1
		}
		break
	}
	return next
}

func NewArchetypeId[T []any](componentTypes T) uint32 {
	id := componentsToHash(componentTypes)
	return id
}

// componentLocator stores a pointer to the archetype that hold the entity and its index in the slice
type componentLocator struct {
	Archetype *Archetype
	Location  int
}

// EntityComponentStore componentStore maps an entity to an array of indices corresponding to the location
type EntityComponentStore struct {
	Entities map[Entity]componentLocator
}

func componentsToHash(components ...interface{}) uint32 {
	h := fnv.New32()
	var sum uint32 = 0
	for _, v := range components {
		h.Reset()
		name := []byte(reflect.TypeOf(v).Name())
		h.Write(name)
		sum += h.Sum32()
	}

	return sum
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
