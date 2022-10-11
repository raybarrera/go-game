package ecs

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {

	t.Run("some test", func(t *testing.T) {
		f := func(i, j int) { print("s:") }
		e := Foreach(f)
		if e != nil {
			t.Error(e)
		}
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
