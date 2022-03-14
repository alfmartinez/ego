package display

import (
	"ego/pkg/configuration"
	"image"
	"image/color"
	"testing"
)

type FakeSprite struct {
	position image.Point
}

func (f *FakeSprite) Doing() string         { return "fake" }
func (f *FakeSprite) Name() string          { return "fake" }
func (f *FakeSprite) Path() string          { return "fake.png" }
func (f *FakeSprite) Position() image.Point { return f.position }
func (f *FakeSprite) Size() uint            { return 1 }

type FakeLoader struct {
	spriteFunc func(string, uint) image.Image
}

func (l *FakeLoader) Init() {}

func (l *FakeLoader) GetSprite(s string, i uint) image.Image {
	return l.spriteFunc(s, i)
}

func TestCreateDisplayCanBuildRts(t *testing.T) {
	actual := CreateDisplay(configuration.Display{
		Type: "rts",
	})
	if _, ok := actual.(*rts); !ok {
		t.Errorf("Create Display for rts should return rts")
	}
}

func TestGetSizeReturnsViewport(t *testing.T) {
	sut := &rts{
		loader: &FakeLoader{},
		config: configuration.Display{
			Size:     configuration.Size{Width: 10, Height: 10},
			ViewPort: configuration.Size{Width: 5, Height: 5},
		}}
	sut.Init()
	size := sut.GetSize()
	expectedSize := configuration.Size{Width: 5, Height: 5}
	if size != expectedSize {
		t.Errorf("Expected size %+v, got %+v", expectedSize, size)
	}
}

func TestRenderReturnsBlankImageSameViewPort(t *testing.T) {
	sut := &rts{
		loader: &FakeLoader{},
		config: configuration.Display{
			Size:     configuration.Size{Width: 10, Height: 10},
			ViewPort: configuration.Size{Width: 10, Height: 10},
		}}
	sut.Init()
	actual := sut.Render()
	if !actual.Bounds().Eq(image.Rect(0, 0, 10, 10)) {
		t.Errorf("Image should be 10x10, got %+v instead", actual.Bounds())
	}
	expectedColor := color.RGBA{0, 0, 255, 255}
	color := actual.At(0, 0)
	if color != expectedColor {
		t.Errorf("Sampled at 0,0; expected %+v, got %+v", expectedColor, color)
	}

}

func TestRenderReturnsBlankImageDifferentViewPort(t *testing.T) {
	sut := &rts{
		loader: &FakeLoader{},
		config: configuration.Display{
			Size:     configuration.Size{Width: 10, Height: 10},
			ViewPort: configuration.Size{Width: 5, Height: 5},
		}}
	sut.Init()
	actual := sut.Render()
	if !actual.Bounds().Eq(image.Rect(0, 0, 5, 5)) {
		t.Errorf("Image should be 10x10, got %+v instead", actual.Bounds())
	}
	expectedColor := color.RGBA{0, 0, 255, 255}
	color := actual.At(0, 0)
	if color != expectedColor {
		t.Errorf("Sampled at 0,0; expected %+v, got %+v", expectedColor, color)
	}

}

func TestRenderSinglePointSprite(t *testing.T) {

	loaderFunc := func(s string, i uint) image.Image {
		if s != "fake.png" {
			t.Errorf("Loader expected path 'fake.png', got %s", s)
		}
		if i != 1 {
			t.Errorf("Loader expected size 1, got %d", i)
		}
		sprite := image.NewRGBA(image.Rect(0, 0, 1, 1))
		sprite.Set(0, 0, color.RGBA{255, 0, 0, 255})
		return sprite
	}

	sut := &rts{
		loader: &FakeLoader{loaderFunc},
		config: configuration.Display{
			Size:     configuration.Size{Width: 10, Height: 10},
			ViewPort: configuration.Size{Width: 10, Height: 10},
		}}
	sprite := &FakeSprite{image.Pt(2, 2)}
	sut.Init()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code panicked")
		}
	}()
	sut.AddObject(sprite)
	actual := sut.Render()
	if !actual.Bounds().Eq(image.Rect(0, 0, 10, 10)) {
		t.Errorf("Image should be 10x10, got %+v instead", actual.Bounds())
	}
	expectedColor := color.RGBA{0, 0, 255, 255}
	bufferColor := actual.At(0, 0)
	if bufferColor != expectedColor {
		t.Errorf("Sampled at 0,0; expected %+v, got %+v", expectedColor, bufferColor)
	}
	spriteColor := actual.At(2, 2)
	expectedSprite := color.RGBA{255, 0, 0, 255}
	if spriteColor != expectedSprite {
		t.Errorf("Sampled at 2,2; expected %+v, got %+v", expectedSprite, spriteColor)
	}
}

func TestAddObjectShouldPanicIfSpriteUnknown(t *testing.T) {

	loaderFunc := func(s string, i uint) image.Image {
		return nil
	}

	sut := &rts{
		loader: &FakeLoader{loaderFunc},
		config: configuration.Display{
			Size:     configuration.Size{Width: 10, Height: 10},
			ViewPort: configuration.Size{Width: 10, Height: 10},
		}}
	sprite := &FakeSprite{image.Pt(2, 2)}
	sut.Init()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	sut.AddObject(sprite)

}
