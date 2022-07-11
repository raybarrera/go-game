package ecs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEntityManager_Queries(t *testing.T) {
	t.Run("Test queries work", func(t *testing.T) {
		sut := &World{
			EntityManager: EntityManager{
				Entities: map[Id][]interface{}{
					NewId(): {""},
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
		for key, val := range result {
			fmt.Println(fmt.Sprintf("%v, %v", key.String(), val))
		}
	})
}

func TestWorld_QueryEntities(t *testing.T) {
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
