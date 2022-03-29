package text

import (
	"os"
)

func ExampleStory_Simple() {
	filepath := "testdata/simple.txt"
	story := CreateStory(filepath)
	story.SetWriter(os.Stdout)
	story.Test()

	// Output:
	// The Simple Story
	//
	// Black Room
	// You see a room with walls black as coal.
	//
	// > east
	//
	// White Room
	// You see a room with its walls white as milk.
	//
}
