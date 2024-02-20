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

// String returns the string representation of the entity.
func (id Entity) String() string {
	return uuid.UUID(id).String()
}

func NewEntityID() Entity {
	id, _ := uuid.NewUUID()
	return Entity(id)
}

// SystemUpdater processes an update/logic on a given collection of components
type SystemUpdater interface {
	Update(deltaTime float64)
}

// World manages all systems and entities
type World struct {
	SystemUpdaters    []SystemUpdater
	ArchetypeDatabase map[uint32]Archetype
}

func NewWorld() *World {
	return &World{
		SystemUpdaters:    make([]SystemUpdater, 0, 10),
		ArchetypeDatabase: make(map[uint32]Archetype),
	}
}

// AddSystem adds a system for the given world to manage.
func (w *World) AddSystem(system SystemUpdater) {
	w.SystemUpdaters = append(w.SystemUpdaters, system)
}

// An Entity is essentially a collection of components.
// Entities are stored in an ArchetypeStore which maps a hash of the components to an Archetype.
func (w *World) CreateEntity(components []any) {
	h := createComponentHash(components...)
	arch, ok := w.getEntityFromStore(components...)
	if !ok {
		arch = createNewArchetype(components...)
	}
	arch.AddEntity(components)
	w.ArchetypeDatabase[h] = *arch
	fmt.Println(arch.PrettyPrint())
}

// createNewArchetype creates a new archetype with the given components.
func createNewArchetype(components ...any) *Archetype {
	h := createComponentHash(components...)
	var types []reflect.Type
	for _, v := range components {
		types = append(types, reflect.TypeOf(v))
	}
	arch := NewEmptyArchetype(h, types...)
	return arch
}

func (w *World) getEntityFromStore(components ...any) (*Archetype, bool) {
	h := createComponentHash(components...)
	if val, ok := w.ArchetypeDatabase[h]; ok {
		return &val, true
	}
	return nil, false
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
	ID             uint32
	NextIndex      int
	componentTable map[reflect.Type][]any
}

// Create a new empty archetype for the given component types and id.
func NewEmptyArchetype(id uint32, componentTypes ...reflect.Type) *Archetype {
	archetype := Archetype{
		ID:             id,
		NextIndex:      0,
		componentTable: make(map[reflect.Type][]any),
	}
	for _, componentType := range componentTypes {
		archetype.componentTable[componentType] = make([]any, 1)
	}
	return &archetype
}

// Create a function that pretty prints the archetype.
func (a *Archetype) PrettyPrint() string {
	var sb strings.Builder

	sb.WriteString("Archetype:\n")
	sb.WriteString(fmt.Sprintf("  Id: %d\n", a.ID))
	sb.WriteString(fmt.Sprintf("  NextIndex: %d\n", a.NextIndex))
	sb.WriteString("  ComponentGroup:\n")

	for t, group := range a.componentTable {
		sb.WriteString(fmt.Sprintf("    %s:\n", t.String()))
		for i, component := range group {
			sb.WriteString(fmt.Sprintf("      - [%d] %v\n", i, component))
		}
	}

	return sb.String()
}

func (a *Archetype) String() string {
	s := ""
	s += fmt.Sprintf("Archetype ID: %v\n", a.ID)
	for k, v := range a.componentTable {
		s += fmt.Sprintf("| type: %v | items: %v |\n", k.String(), len(v))
	}
	s += fmt.Sprintf("| NextIndex: %v | Valid: %v |\n", a.NextIndex, a.GetNextAvailableIndex())
	return s
}

func (archetype *Archetype) AddEntity(components []any) error {
	appendMode := false
	if isFull := archetype.GetNextAvailableIndex() == -1; isFull {
		appendMode = true
	}
	for _, v := range components {
		t := reflect.TypeOf(v)
		if _, ok := archetype.componentTable[t]; !ok {
			return fmt.Errorf("component type %v not found in archetype", t)
		}
		if appendMode {
			archetype.componentTable[t] = append(archetype.componentTable[t], v)
		} else {
			archetype.componentTable[t][archetype.NextIndex] = v
		}

	}
	archetype.NextIndex++
	return nil
}

// GetNextAvailableIndex returns the next available index in the archetype.
// If the archetype is full, it returns -1.
func (a *Archetype) GetNextAvailableIndex() int {
	if isIndexAvailable(a.componentTable, a.NextIndex) {
		return a.NextIndex
	}
	next := a.NextIndex
	for _, componentSlice := range a.componentTable {
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

func (a *Archetype) GetNextAvailableIndexOptimized() int {
	if isIndexAvailable(a.componentTable, a.NextIndex) {
		return a.NextIndex
	}
	shortestSlice := -1
	for _, componentSlice := range a.componentTable {
		if shortestSlice == -1 || len(componentSlice) < shortestSlice {
			shortestSlice = len(componentSlice)
		}
	}

	if shortestSlice <= 0 {
		return 0
	}

	for i := 0; i < shortestSlice; i++ {
		availableInAll := true
		for _, componentSlice := range a.componentTable {
			if componentSlice[i] != nil {
				availableInAll = false

				break
			}
		}

		if availableInAll {
			return i
		}
	}
	return shortestSlice
}

func isIndexAvailable(componentTable map[reflect.Type][]any, index int) bool {
	for _, componentSlice := range componentTable {
		if len(componentSlice) <= index {
			return false
		}
		if componentSlice[index] != nil {
			return false
		}
	}
	return true
}

func NewArchetypeId[T []any](componentTypes T) uint32 {
	id := createComponentHash(componentTypes)
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

func createComponentHash(components ...interface{}) uint32 {
	h := fnv.New32()
	var sum uint32 = 0
	for _, component := range components {
		h.Reset()
		name := []byte(reflect.TypeOf(component).Name())
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
