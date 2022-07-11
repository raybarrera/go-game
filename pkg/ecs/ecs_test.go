package ecs

import (
	"reflect"
	"testing"
)

func TestWorld_EntityManager_Queries(t *testing.T) {
	t.Run("Test queries work", func(t *testing.T) {
		id := NewId()
		expectedString := "abc"
		expectedLength := 1
		sut := &World{
			EntityManager: EntityManager{
				Entities: map[Id][]interface{}{
					id: {expectedString},
				},
			},
		}
		result, err := sut.QueryEntities(reflect.TypeOf(""))
		if err != nil {
			t.Errorf(err.Error())
		}
		if result == nil {
			t.Errorf("Something")
		}
		if result[id] == nil {
			t.Errorf("ID not found in from query")
		}
		if len(result[id]) < expectedLength {
			t.Errorf("Incorrect length in the result array")
		}
		if result[id][0] != expectedString {
			t.Errorf("Wanted \"%v\" in the 0 position of the result, got \"%v\"", expectedString, result[id][0])
		}
	})
	t.Run("Test multiple matching entities in query", func(t *testing.T) {
		sut := &World{
			EntityManager: EntityManager{
				Entities: map[Id][]interface{}{
					NewId(): {""},
					NewId(): {2.0},
					NewId(): {3.0},
				},
			},
		}
		result, err := sut.QueryEntities(reflect.TypeOf(0.0))
		if err != nil {
			t.Errorf(err.Error())
		}
		if result == nil {
			t.Errorf("Nil result")
		}
		if len(result) != 2 {
			t.Errorf("Incorrect length in the result array, wanted %v, got %v", 2, len(result))
		}
	})
}

func TestWorld_QueryEntities_OLD(t *testing.T) {
	type args struct {
		components []reflect.Type
	}
	tests := []struct {
		name    string
		w       *World
		args    args
		want    []interface{}
		wantErr bool
	}{
		{
			name: "Check we get the right types back",
			w: &World{
				EntityManager: EntityManager{
					Entities: map[Id][]interface{}{
						NewId(): {""},
					},
				},
			},
			args: args{
				components: []reflect.Type{
					reflect.TypeOf(""),
				},
			},
			want: []interface{}{
				&args{},
			},
			wantErr: false,
		},
		{
			name: "Check multiple entities",
			w: &World{
				EntityManager: EntityManager{
					Entities: map[Id][]interface{}{
						NewId(): {""},
						NewId(): {"", 2.0},
					},
				},
			},
			args: args{
				components: []reflect.Type{
					reflect.TypeOf(""),
				},
			},
			want: []interface{}{
				&args{},
				1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.w.QueryEntities(tt.args.components...)
			if (err != nil) != tt.wantErr {
				t.Errorf("World.QueryEntities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("World.QueryEntities() = %v, want %v", got, tt.want)
			}
		})
	}
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
