package movement

import (
	"github.com/alfmartinez/ego/engine/physics/matrix"
	"reflect"
	"testing"
	"time"
)

func TestCreateMovementReturnsMovement(t *testing.T) {
	i := CreateMovement(Location{1, 10})

	if i.Position().X != 1 {
		t.Errorf("should be initialized with X value")
	}

	if i.Position().Y != 10 {
		t.Errorf("should be initialized with Y value")
	}
}

func Test_movement_GetMatrix(t *testing.T) {
	type fields struct {
		matrix matrix.M
	}
	tests := []struct {
		name   string
		fields fields
		want   matrix.M
	}{
		{
			"sample",
			fields{
				matrix.M{
					P: matrix.Vec2{X: 1, Y: 0},
					S: matrix.Vec2{X: 0, Y: 1},
					A: matrix.Vec2{X: 1, Y: 1},
				},
			},
			matrix.M{
				P: matrix.Vec2{X: 1, Y: 0},
				S: matrix.Vec2{X: 0, Y: 1},
				A: matrix.Vec2{X: 1, Y: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"Constructor", func(t *testing.T) {
			m := &movement{
				matrix: tt.fields.matrix,
			}
			if got := m.GetMatrix(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("movement.GetMatrix() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name+"Setter", func(t *testing.T) {
			m := &movement{matrix: matrix.M{}}
			m.SetMatrix(tt.fields.matrix)
			if got := m.GetMatrix(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("movement.GetMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_movement_MoveDirection(t *testing.T) {
	type fields struct {
		matrix matrix.M
	}
	type args struct {
		direction Direction
		dt        time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   matrix.M
	}{
		{
			"0,0 - UP",
			fields{
				matrix.M{
					P: matrix.Vec2{},
					S: matrix.Vec2{},
					A: matrix.Vec2{},
				},
			},
			args{
				MOVE_UP,
				time.Second,
			},
			matrix.M{
				P: matrix.Vec2{},
				S: matrix.Vec2{X: 0, Y: -50},
				A: matrix.Vec2{},
			},
		},
		{
			"0,0 - DOWN",
			fields{
				matrix.M{
					P: matrix.Vec2{},
					S: matrix.Vec2{},
					A: matrix.Vec2{},
				},
			},
			args{
				MOVE_DOWN,
				time.Second,
			},
			matrix.M{
				P: matrix.Vec2{},
				S: matrix.Vec2{X: 0, Y: 50},
				A: matrix.Vec2{},
			},
		},
		{
			"0,0 - RIGHT",
			fields{
				matrix.M{
					P: matrix.Vec2{},
					S: matrix.Vec2{},
					A: matrix.Vec2{},
				},
			},
			args{
				MOVE_RIGHT,
				time.Second,
			},
			matrix.M{
				P: matrix.Vec2{},
				S: matrix.Vec2{X: 50, Y: 0},
				A: matrix.Vec2{},
			},
		},
		{
			"0,0 - LEFT",
			fields{
				matrix.M{
					P: matrix.Vec2{},
					S: matrix.Vec2{},
					A: matrix.Vec2{},
				},
			},
			args{
				MOVE_LEFT,
				time.Second,
			},
			matrix.M{
				P: matrix.Vec2{},
				S: matrix.Vec2{X: -50, Y: 0},
				A: matrix.Vec2{},
			},
		},
		{
			"0,0 - UP, reverse",
			fields{
				matrix.M{
					P: matrix.Vec2{},
					S: matrix.Vec2{},
					A: matrix.Vec2{},
				},
			},
			args{
				MOVE_UP,
				-time.Second,
			},
			matrix.M{
				P: matrix.Vec2{},
				S: matrix.Vec2{X: 0, Y: 50},
				A: matrix.Vec2{},
			},
		},
		{
			"0,0 - DOWN, reserve",
			fields{
				matrix.M{
					P: matrix.Vec2{},
					S: matrix.Vec2{},
					A: matrix.Vec2{},
				},
			},
			args{
				MOVE_DOWN,
				-time.Second,
			},
			matrix.M{
				P: matrix.Vec2{},
				S: matrix.Vec2{X: 0, Y: -50},
				A: matrix.Vec2{},
			},
		},
		{
			"0,0 - RIGHT, reverse",
			fields{
				matrix.M{
					P: matrix.Vec2{},
					S: matrix.Vec2{},
					A: matrix.Vec2{},
				},
			},
			args{
				MOVE_RIGHT,
				-time.Second,
			},
			matrix.M{
				P: matrix.Vec2{},
				S: matrix.Vec2{X: -50, Y: 0},
				A: matrix.Vec2{},
			},
		},
		{
			"0,0 - LEFT, reverse",
			fields{
				matrix.M{
					P: matrix.Vec2{},
					S: matrix.Vec2{},
					A: matrix.Vec2{},
				},
			},
			args{
				MOVE_LEFT,
				-time.Second,
			},
			matrix.M{
				P: matrix.Vec2{},
				S: matrix.Vec2{X: 50, Y: 0},
				A: matrix.Vec2{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &movement{
				matrix: tt.fields.matrix,
			}
			m.MoveDirection(tt.args.direction, tt.args.dt)
			if got := m.GetMatrix(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("movement.GetMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
