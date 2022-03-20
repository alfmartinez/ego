package matrix

import (
	"reflect"
	"testing"
)

func TestVec2_Add(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		d Vec2
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec2
	}{
		{
			"Null add",
			fields{0, 0},
			args{Vec2{0, 0}},
			Vec2{},
		},
		{
			"Null add, 1-1",
			fields{1, 1},
			args{Vec2{0, 0}},
			Vec2{1, 1},
		},
		{
			"1-2 add, 1-1",
			fields{1, 1},
			args{Vec2{1, 2}},
			Vec2{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec2{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Add(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec2.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Mul(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		f float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec2
	}{
		{
			"Zero",
			fields{1, 1},
			args{0},
			Vec2{},
		},
		{
			"One",
			fields{10, 5},
			args{1},
			Vec2{10, 5},
		},
		{
			"Other",
			fields{2, 3},
			args{10},
			Vec2{20, 30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec2{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Mul(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec2.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}
