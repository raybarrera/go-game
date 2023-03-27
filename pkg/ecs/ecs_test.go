package ecs

import (
	"reflect"
	"testing"
)

// TestWorld_CreateEntity tests the creation of an TestWorld_CreateEntity
func TestWorld_CreateEntity(t *testing.T) {
	w := NewWorld()
	w.CreateEntity([]interface{}{1})
}

// Test that ints and floats produce different hashes
func TestHash_IntAndFloat(t *testing.T) {
	a := componentsToHash(1)
	b := componentsToHash(1.0)
	if ok := a != b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

// Test that maps produce matching hashes when the keys and values are the same
func TestHash_Map(t *testing.T) {
	a := componentsToHash(map[string]int{"a": 1})
	b := componentsToHash(map[string]int{"a": 1})
	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

// Test that maps with same keys and different values produce the same hashes
func TestHash_MapWithDifferentValues(t *testing.T) {
	a := componentsToHash(map[string]int{"a": 1})
	b := componentsToHash(map[string]int{"a": 2})
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
	a := componentsToHash(A{1})
	b := componentsToHash(B{A{1}, 2})
	if ok := a != b; !ok {
		t.Errorf("Expected %v == %v got %v \n", a, b, ok)
	}
}

func TeshHash_IsDeterministic(t *testing.T) {
	a := componentsToHash("s")
	b := componentsToHash("s")

	if ok := int(a) == int(b); !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_HashingMultipleTimesReturnsSameResult(t *testing.T) {
	a := componentsToHash("s")
	b := componentsToHash("s")
	b = componentsToHash("s")

	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfSameType(t *testing.T) {
	a := componentsToHash("s", "D")
	b := componentsToHash("s", "D")
	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfDifferentType(t *testing.T) {
	a := componentsToHash("s", 1)
	b := componentsToHash("s", 1)
	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfDifferentTypeDifferentOrder(t *testing.T) {
	a := componentsToHash("s", 1)
	b := componentsToHash(1, "s")
	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfDifferentTypesMismatch(t *testing.T) {
	a := componentsToHash("s", 1.5)
	b := componentsToHash(1, "s")
	if ok := a != b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_MismatchedArgumentCountBreaks(t *testing.T) {
	a := componentsToHash("s", 1, "d")
	b := componentsToHash(1, "s")
	if ok := a != b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfSameTypeWithDiffrerentOrder(t *testing.T) {
	a := componentsToHash("s", "D")
	b := componentsToHash("D", "s")
	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfSameTypeWithDiffrerentOrderInverse(t *testing.T) {
	a := componentsToHash("s", "d")
	b := componentsToHash("D", "s")
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
	actual := a.GetNextIndex()

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetNextIndexId_NoAvailableSlots(t *testing.T) {
	a := &Archetype{
		componentGroup: map[reflect.Type][]any{
			reflect.TypeOf(&ComponentData{}): {
				&ComponentData{}, &ComponentData{},
			},
			reflect.TypeOf(&OtherComponent{}): {
				&OtherComponent{}, &OtherComponent{}, &OtherComponent{},
			},
		},
	}

	expected := -1
	actual := a.GetNextIndex()
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetNextIndexId_SomeAvailableSlots(t *testing.T) {
	a := &Archetype{
		componentGroup: map[reflect.Type][]any{
			reflect.TypeOf(&ComponentData{}): {
				&ComponentData{}, nil,
			},
			reflect.TypeOf(&OtherComponent{}): {
				&OtherComponent{}, OtherComponent{}, nil,
			},
		},
	}

	expected := 1
	actual := a.GetNextIndex()
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestFindNextAvailableIndex_OneAvailable(t *testing.T) {
	a := &Archetype{
		componentGroup: map[reflect.Type][]any{
			reflect.TypeOf(&ComponentData{}): {
				&ComponentData{}, nil,
			},
			reflect.TypeOf(&OtherComponent{}): {
				&OtherComponent{}, nil,
			},
		},
	}

	expected := 1
	actual := a.FindNextAvailableIndex()
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
