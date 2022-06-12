package tiled

import (
	"reflect"
	"testing"
)

func TestParseMapString(t *testing.T) {
	type args struct {
		mapString string
	}
	tests := []struct {
		name    string
		args    args
		want    TileMap
		wantErr bool
	}{
		{
			name:    "Empty array",
			args:    args{mapString: "[]"},
			want:    TileMap{},
			wantErr: true,
		},
		{
			name:    "valid with width 100",
			args:    args{mapString: "{\"width\":100}"},
			want:    TileMap{Width: 100},
			wantErr: false,
		},
		{
			name:    "valid with incorrect capitalization on width",
			args:    args{mapString: "{\"Width\":100}"},
			want:    TileMap{Width: 100},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMapString(tt.args.mapString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMapString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMapString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
