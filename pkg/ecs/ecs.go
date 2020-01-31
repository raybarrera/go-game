package ecs

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
	Update(dt float64)
}

// Start does nothing currently
func (*Entity) Start() {

}

// AddComponent adds a component to an entity
func (e *Entity) AddComponent(c Component) error {
	e.components = append(e.components, c)
	return nil
}
