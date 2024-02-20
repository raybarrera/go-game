package ecs

import (
	"reflect"
	"testing"
	"time"
)

// TestWorld_CreateEntity tests the creation of an TestWorld_CreateEntity
func TestWorld_CreateEntity(t *testing.T) {
	w := NewWorld()
	w.CreateEntity([]interface{}{1})
}

// Test that ints and floats produce different hashes
func TestHash_IntAndFloat(t *testing.T) {
	a := createComponentHash(1)
	b := createComponentHash(1.0)
	if ok := a != b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

// Test that maps produce matching hashes when the keys and values are the same
func TestHash_Map(t *testing.T) {
	a := createComponentHash(map[string]int{"a": 1})
	b := createComponentHash(map[string]int{"a": 1})
	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

// Test that maps with same keys and different values produce the same hashes
func TestHash_MapWithDifferentValues(t *testing.T) {
	a := createComponentHash(map[string]int{"a": 1})
	b := createComponentHash(map[string]int{"a": 2})
	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

// Test that a derived type produces a different hash than the base type
func TestHash_DerivedType(t *testing.T) {
	type A struct {
		a int
	}
	type B struct {
		A
		b int
	}
	a := createComponentHash(A{1})
	b := createComponentHash(B{A{1}, 2})
	if ok := a != b; !ok {
		t.Errorf("Expected %v == %v got %v \n", a, b, ok)
	}
}

func TeshHash_IsDeterministic(t *testing.T) {
	a := createComponentHash("s")
	b := createComponentHash("s")

	if ok := int(a) == int(b); !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_HashingMultipleTimesReturnsSameResult(t *testing.T) {
	a := createComponentHash("s")
	b := createComponentHash("s")
	b = createComponentHash("s")

	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfSameType(t *testing.T) {
	a := createComponentHash("s", "D")
	b := createComponentHash("s", "D")
	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfDifferentType(t *testing.T) {
	a := createComponentHash("s", 1)
	b := createComponentHash("s", 1)
	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfDifferentTypeDifferentOrder(t *testing.T) {
	a := createComponentHash("s", 1)
	b := createComponentHash(1, "s")
	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfDifferentTypesMismatch(t *testing.T) {
	a := createComponentHash("s", 1.5)
	b := createComponentHash(1, "s")
	if ok := a != b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_MismatchedArgumentCountBreaks(t *testing.T) {
	a := createComponentHash("s", 1, "d")
	b := createComponentHash(1, "s")
	if ok := a != b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfSameTypeWithDiffrerentOrder(t *testing.T) {
	a := createComponentHash("s", "D")
	b := createComponentHash("D", "s")
	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfSameTypeWithDiffrerentOrderInverse(t *testing.T) {
	a := createComponentHash("s", "d")
	b := createComponentHash("D", "s")
	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func Test(t *testing.T) {
	t.Run("some test", func(t *testing.T) {
		w := NewWorld()
		w.CreateEntity([]interface{}{1})
	})
}

func TestContainsType(t *testing.T) {
	type args struct {
		arr   []interface{}
		check reflect.Type
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test simple case, give string, get string",
			args: args{
				arr:   []interface{}{""},
				check: reflect.TypeOf(""),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsType(tt.args.arr, tt.args.check); got != tt.want {
				t.Errorf("ContainsType() = %v, want %v", got, tt.want)
			}
		})
	}
}

type ComponentData struct{}
type OtherComponent struct{}

func TestGetNextIndexId(t *testing.T) {
	a := &Archetype{}

	expected := 0
	actual := a.GetNextAvailableIndex()

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetNextIndexId_NoAvailableSlots(t *testing.T) {
	a := &Archetype{
		componentTable: map[reflect.Type][]any{
			reflect.TypeOf(&ComponentData{}): {
				&ComponentData{}, &ComponentData{},
			},
			reflect.TypeOf(&OtherComponent{}): {
				&OtherComponent{}, &OtherComponent{}, &OtherComponent{},
			},
		},
	}

	expected := -1
	actual := a.GetNextAvailableIndex()
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetNextIndexId_SomeAvailableSlots(t *testing.T) {
	a := &Archetype{
		componentTable: map[reflect.Type][]any{
			reflect.TypeOf(&ComponentData{}): {
				&ComponentData{}, nil,
			},
			reflect.TypeOf(&OtherComponent{}): {
				&OtherComponent{}, OtherComponent{}, nil,
			},
		},
	}

	expected := 1
	actual := a.GetNextAvailableIndex()
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

// Define a complex archetype variable with a large componentTable containing slices with 20 components for benchmarking
var complexArchetype = &Archetype{
	componentTable: map[reflect.Type][]any{
		reflect.TypeOf(&ComponentData{}): {
			&ComponentData{}, nil, &ComponentData{}, nil, &ComponentData{}, nil, &ComponentData{}, nil, &ComponentData{}, nil,
			&ComponentData{}, nil, &ComponentData{}, nil, &ComponentData{}, nil, &ComponentData{}, nil, &ComponentData{}, nil,
		},
		reflect.TypeOf(&OtherComponent{}): {
			&OtherComponent{}, nil, &OtherComponent{}, nil, &OtherComponent{}, nil, &OtherComponent{}, nil, &OtherComponent{}, nil,
			&OtherComponent{}, nil, &OtherComponent{}, nil, &OtherComponent{}, nil, &OtherComponent{}, nil, &OtherComponent{}, nil,
		},
	},
}

func BenchmarkAllGetNextAvailableIndex(b *testing.B) {
	b.Run("Normal", func(b *testing.B) {
		a := *complexArchetype
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			a.GetNextAvailableIndex()
		}
	})
	b.Run("Optimized Miss", func(b *testing.B) {
		a := *complexArchetype
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			a.GetNextAvailableIndexOptimized()
		}
	})
	b.Run("Optimized Hit", func(b *testing.B) {
		a := *complexArchetype
		a.NextIndex = 1
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			a.GetNextAvailableIndexOptimized()
		}
	})
}

func measureFuncTime(f func() int) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}

func measureTime(f func(...interface{}) uint32, components ...interface{}) time.Duration {
	start := time.Now()
	f(components...)
	return time.Since(start)
}

type Behavior interface {
	Act() string
}
type PolymorphicComponent struct {
	Behaviors []Behavior
}
type SpecificBehavior struct {
	Description string
}

func (sb SpecificBehavior) Act() string {
	return sb.Description
}
func BenchmarkComponentsToHash(b *testing.B) {
	type NestedComponent struct {
		Timestamp time.Time
		Data      []byte
	}

	simpleStruct := struct{ name string }{name: "Simple"}
	embeddedStruct := struct{ Simple, additionalField string }{Simple: "Embedded", additionalField: "Data"}
	complexStruct := struct {
		Coordinates []float64
		Properties  map[string]string
		Nested      NestedComponent
	}{
		Coordinates: []float64{1.0, 2.0},
		Properties:  map[string]string{"a": "b"},
		Nested:      NestedComponent{Timestamp: time.Now(), Data: []byte("data")},
	}
	polymorphicComponent := PolymorphicComponent{
		Behaviors: []Behavior{
			SpecificBehavior{Description: "Behavior 1"},
			SpecificBehavior{Description: "Behavior 2"},
		},
	}
	b.Run("Simple", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			createComponentHash(simpleStruct)
		}
	})

	b.Run("Embedded", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			createComponentHash(embeddedStruct)
		}
	})
	b.Run("Complex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			createComponentHash(complexStruct)
		}
	})
	b.Run("Polymorphic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			createComponentHash(polymorphicComponent)
		}
	})
}
