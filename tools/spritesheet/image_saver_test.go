package spritesheet

import (
	"image"
	"image/color"
	_ "image/png"
	"os"
	"testing"
)

func TestSaveImage(t *testing.T) {
	pwd, _ := os.Getwd()
	defer func() {
		os.Remove(pwd + "/testdata/pixel.png")
	}()
	img := image.NewNRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.NRGBA{0, 128, 128, 128})

	type args struct {
		path string
		img  image.Image
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		panic   bool
	}{
		{
			"save",
			args{
				"/testdata/pixel.png",
				img,
			},
			false,
			false,
		},
		{
			"save mssing directory",
			args{
				"/testdata/foo/pixel.png",
				img,
			},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					r := recover()
					if r == nil {
						t.Error("should panic")
					}
				}()
			}
			if err := SaveImage(tt.args.path, tt.args.img); (err != nil) != tt.wantErr {
				t.Errorf("SaveImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
