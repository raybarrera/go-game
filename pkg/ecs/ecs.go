package ecs

type Entity struct {
	components []Component
}

type Component struct {
}

type Ticker interface {
	onUpdate(dt float32)
}
