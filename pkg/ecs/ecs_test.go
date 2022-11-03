package ecs

import (
	"reflect"
	"testing"
)

func TeshHash_IsDeterministic(t *testing.T) {
	a := hash("s")
	b := hash("s")

	if ok := int(a) == int(b); !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_HashingMultipleTimesReturnsSameResult(t *testing.T) {
	a := hash("s")
	b := hash("s")
	b = hash("s")

	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfSameType(t *testing.T) {
	a := hash("s", "D")
	b := hash("s", "D")
	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfDifferentType(t *testing.T) {
	a := hash("s", 1)
	b := hash("s", 1)
	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfDifferentTypeDifferentOrder(t *testing.T) {
	a := hash("s", 1)
	b := hash(1, "s")
	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfDifferentTypesMismatch(t *testing.T) {
	a := hash("s", 1.5)
	b := hash(1, "s")
	if ok := a != b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_MismatchedArgumentCountBreaks(t *testing.T) {
	a := hash("s", 1, "d")
	b := hash(1, "s")
	if ok := a != b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}

func TestHash_UsingMultipleInputsOfSameTypeWithDiffrerentOrder(t *testing.T) {
	a := hash("s", "D")
	b := hash("D", "s")
	if ok := a == b; !ok {
		t.Errorf("Expected %v, got %v \n", true, ok)
	}
}
func TestHash_UsingMultipleInputsOfSameTypeWithDiffrerentOrderInverse(t *testing.T) {
	a := hash("s", "d")
	b := hash("D", "s")
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
