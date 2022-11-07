package main

import (
	"fmt"
	"go-game/pkg/ecs"
)

var world *ecs.World

func main() {
	world = ecs.NewWorld()
	world.CreateEntity([]any{1})
	world.CreateEntity([]any{1, "something"})
	world.CreateEntity([]any{func() {}, 1, &SystemComponents{}})
	RunTest()
}

func RunTest() {

}

type SystemComponents struct {
	Health int
	Damage int
}

var container SystemComponents

func Update(dt float64) {
	var db map[string][]interface{} = make(map[string][]interface{})
	db["key"] = []interface{}{1}
	db["otherKey"] = []interface{}{"1", "2"}
	fmt.Printf("key: %v\n otherkey: %v", db["key"], db["otherKey"])

}
