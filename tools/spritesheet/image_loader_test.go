package spritesheet

import (
	"image"
	"image/color"
	_ "image/png"
	"os"
	"reflect"
	"testing"
)

func TestLoadImage(t *testing.T) {
	pwd, _ := os.Getwd()

	defer func() {
		os.Remove(pwd + "/testdata/load.png")
	}()
	img := image.NewNRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.NRGBA{255, 128, 128, 128})
	SaveImage("/testdata/load.png", img)

	type args struct {
		path string
	}
	tests := []struct {
		name  string
		args  args
		want  image.Image
		panic bool
	}{
		{
			"missing",
			args{
				pwd + "/testdata/missing.png",
			},
			nil,
			true,
		},
		{
			"missing",
			args{
				pwd + "/testdata/load.png",
			},
			img,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if tt.panic && r == nil {
					t.Error("should panic")
				}
				if !tt.panic && r != nil {
					t.Error(r)
				}
			}()
			if got := LoadImage(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
