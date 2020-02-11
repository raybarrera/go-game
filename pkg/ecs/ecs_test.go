package ecs

import (
	"reflect"
	"testing"
)

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
				Entities: map[interface{}][]reflect.Type{
					&args{}: []reflect.Type{
						reflect.TypeOf(""),
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
				Entities: map[interface{}][]reflect.Type{
					&args{}: []reflect.Type{
						reflect.TypeOf(""),
					},
					1: []reflect.Type{
						reflect.TypeOf(""),
						reflect.TypeOf(2.0),
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
