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
