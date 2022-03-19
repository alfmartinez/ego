package display

import (
	"ego/engine/layer"
	"image"
	"reflect"
	"testing"
)

func TestCreateDisplayable(t *testing.T) {
	type args struct {
		path     string
		size     uint
		position image.Point
		layer    layer.Layer
	}
	tests := []struct {
		name string
		args args
		want Displayable
	}{
		{
			"simple case",
			args{
				"foo",
				1,
				image.Pt(1, 1),
				layer.BACKGROUND,
			},
			displayable{
				"foo",
				1,
				image.Pt(1, 1),
				layer.BACKGROUND,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateDisplayable(tt.args.path, tt.args.size, tt.args.position, tt.args.layer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDisplayable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_displayable_Methods(t *testing.T) {
	type fields struct {
		path     string
		size     uint
		position image.Point
		layer    layer.Layer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"Methods",
			fields{
				"foo",
				2,
				image.Pt(4, 2),
				layer.FAR_BACKGROUND,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := displayable{
				path:     tt.fields.path,
				size:     tt.fields.size,
				position: tt.fields.position,
				layer:    tt.fields.layer,
			}
			if got := d.Path(); got != tt.fields.path {
				t.Errorf("displayable.Path() = %v, want %v", got, tt.fields.path)
			}
			if got := d.Size(); got != tt.fields.size {
				t.Errorf("displayable.Size() = %v, want %v", got, tt.fields.size)
			}
			if got := d.Position(); got != tt.fields.position {
				t.Errorf("displayable.Position() = %v, want %v", got, tt.fields.position)
			}
			if got := d.Layer(); got != tt.fields.layer {
				t.Errorf("displayable.Layer() = %v, want %v", got, tt.fields.layer)
			}
		})
	}
}
