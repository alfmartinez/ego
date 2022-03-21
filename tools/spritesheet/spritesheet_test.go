package spritesheet

import (
	"image"
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	if _, ok := s.(*spritesheet); !ok {
		t.Error("New should return *spritesheet")
	}
}

func Test_spritesheet_Load(t *testing.T) {
	pwd, _ := os.Getwd()
	var mockLoad func(string) image.Image
	var mockSave func(string, image.Image) error
	type fields struct {
		dest      image.Image
		loadImage func(string) image.Image
		saveImage func(string, image.Image) error
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		paths  []string
	}{
		{
			"simple",
			fields{
				loadImage: func(path string) image.Image {
					return mockLoad(path)
				},
				saveImage: func(path string, img image.Image) error {
					return mockSave(path, img)
				},
			},
			args{
				"/testdata/sample/",
			},
			[]string{pwd + "/testdata/sample/.keep", pwd + "/testdata/sample/malformed"},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			paths := make([]string, 0)
			mockLoad = func(path string) image.Image {
				paths = append(paths, path)
				return image.NewNRGBA(image.Rect(0, 0, 10, 10))
			}
			mockSave = func(path string, img image.Image) error {
				return nil
			}
			s := &spritesheet{
				dest:      tt.fields.dest,
				loadImage: tt.fields.loadImage,
				saveImage: tt.fields.saveImage,
			}
			s.Load(tt.args.path)
			s.Export("test")
			if !reflect.DeepEqual(tt.paths, paths) {
				t.Errorf("Expected %v, got %v", tt.paths, paths)
			}
		})
	}
}

func TestRegression(t *testing.T) {
	pwd, _ := os.Getwd()

	defer func() {
		os.Remove(pwd + "/testdata/regression/output/got.png")
	}()

	s := New()
	s.Load("/testdata/regression/input/")
	s.Export("/testdata/regression/output/got.png")

	got := LoadImage(pwd + "/testdata/regression/output/got.png")
	want := LoadImage(pwd + "/testdata/regression/output/sheet.png")

	if !reflect.DeepEqual(got, want) {
		t.Error("images should be equal")
	}

}
