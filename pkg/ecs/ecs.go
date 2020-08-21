package ecs

type Entity struct {
	id         uint64
	components []Component
}

type Component struct {
}

type Ticker interface {
	onUpdate(dt float32)
}

func (*Entity) Start() {

}

func (e *Entity) AddComponent(c Component) error {
	e.components = append(e.components, c)
	return nil
}
