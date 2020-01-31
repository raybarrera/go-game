package ecs

// Entity is a collection of components
type Entity struct {
	id         uint64
	components []Component
}

// Component is a data container
type Component struct {
}

// Ticker does nothing currently
type Ticker interface {
	onUpdate(dt float32)
}

// Start does nothing currently
func (*Entity) Start() {

}

// AddComponent adds a component to an entity
func (e *Entity) AddComponent(c Component) error {
	e.components = append(e.components, c)
	return nil
}
