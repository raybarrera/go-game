package ecs

import (
	"fmt"
	"hash/fnv"
	"reflect"
	"strings"

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
		ArchetypeStore: make(map[uint32]Archetype),
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

// An Entity is essentially a collection of components.
// Entities are stored in an ArchetypeStore which maps a hash of the components to an Archetype.
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
		arch = NewEmptyArchetype(h, types...)
	}
	arch.AddEntity(components)
	w.ArchetypeStore[h] = *arch
	fmt.Println(arch.PrettyPrint())
}

func AddComponent[T any](to Entity, component T) {
	//TODO: Calculate old archetype
	//TODO: Calculate new archetype
	//TODO: Migrate to new archetype
}

// RemoveComponentOfType TODO needs to figure out what happens when we
// have multiple components of the same type
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

// Archetype represents a combination of components. It acts as a store for
// the components of entities that share the same component types.
type Archetype struct {
	Id             uint32
	NextIndex      int
	componentGroup map[reflect.Type][]any
}

// Create a new empty archetype for the given component types and id.
func NewEmptyArchetype(id uint32, componentTypes ...reflect.Type) *Archetype {
	archetype := Archetype{
		Id:             id,
		NextIndex:      0,
		componentGroup: make(map[reflect.Type][]any),
	}
	for _, componentType := range componentTypes {
		archetype.componentGroup[componentType] = make([]any, 1)
	}
	return &archetype
}

// Create a function that pretty prints the archetype.
func (a *Archetype) PrettyPrint() string {
	var sb strings.Builder

	sb.WriteString("Archetype:\n")
	sb.WriteString(fmt.Sprintf("  Id: %d\n", a.Id))
	sb.WriteString(fmt.Sprintf("  NextIndex: %d\n", a.NextIndex))
	sb.WriteString("  ComponentGroup:\n")

	for t, group := range a.componentGroup {
		sb.WriteString(fmt.Sprintf("    %s:\n", t.String()))
		for i, component := range group {
			sb.WriteString(fmt.Sprintf("      - [%d] %v\n", i, component))
		}
	}

	return sb.String()
}

func (a *Archetype) String() string {
	s := ""
	s += fmt.Sprintf("Archetype ID: %v\n", a.Id)
	for k, v := range a.componentGroup {
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
			a.componentGroup[t] = append(a.componentGroup[t], v)
		} else {
			a.componentGroup[t][a.NextIndex] = v
		}

	}
	a.NextIndex++
}

func (a *Archetype) FindNextAvailableIndex() int {
	minLength := -1
	for _, componentSlice := range a.componentGroup {
		if minLength == -1 || len(componentSlice) < minLength {
			minLength = len(componentSlice)
		}
	}

	type result struct {
		index int
		seen  bool
	}

	resChan := make(chan result)

	for i := 0; i < minLength; i++ {
		go func(index int) {
			isNil := false
			for _, componentSlice := range a.componentGroup {
				if componentSlice[index] != nil {
					isNil = false
					break
				}
				isNil = true
			}
			resChan <- result{index, isNil}
		}(i)
	}
	for i := 0; i < minLength; i++ {
		res := <-resChan
		if res.seen {
			a.NextIndex = res.index
			return res.index
		}
	}
	a.NextIndex = -1
	return -1
}

// GetNextIndex returns the next available index in the archetype.
// If the archetype is full, it returns -1.
func (a *Archetype) GetNextIndex() int {
	next := a.NextIndex
	for _, componentSlice := range a.componentGroup {
		isFull := true
		if len(componentSlice) <= 0 {
			isFull = true
			break
		}
		if len(componentSlice) < next && componentSlice[next] == nil {
			break
		}
		for i, componentData := range componentSlice {
			if componentData == nil {
				isFull = false
				next = i
				break
			}
		}
		if isFull {
			a.NextIndex = len(componentSlice)
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
