package main

import (
	"go-game/pkg/ecs"
)

var world *ecs.World

func main() {
	world = ecs.NewWorld()
	world.CreateEntity([]interface{}{1})

	RunTest()
}

func RunTest() {
}
