package text

import (
	"os"
)

func ExampleStory_simple() {
	filepath := "testdata/simple.txt"
	story := CreateStory(filepath)
	story.SetWriter(os.Stdout)
	story.Test()

	// Output:
	// The Simple Story
	//
	// Black Rock
	// You see a rock cut from black stone.
	//
	// > east
	//
	// Sunken Cove
	// You see a room with its walls white as milk.
	//
}

func ExampleStory_bic() {
	filepath := "testdata/bic.txt"
	story := CreateStory(filepath)
	story.SetWriter(os.Stdout)
	story.Test()
	// Output:
	// Bic
	//
	// Staff Break Room
	//
	// > look at orange
	//
	// It's a samll hard pinch-skinned thing from the lunch room, probably with lots of pips and no juice.
	//
	// > look at pen

}
