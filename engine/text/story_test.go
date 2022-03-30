package text

import (
	"os"
)

var paths = []string{
	"data/values.txt",
	"data/kinds.txt",
}

func ExampleStory_simple() {
	filepaths := []string{}
	filepaths = append(filepaths, paths...)
	filepaths = append(filepaths, "testdata/simple.txt")

	story := CreateStory(filepaths, false, false)
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
	filepaths := []string{}
	filepaths = append(filepaths, paths...)
	filepaths = append(filepaths, "testdata/bic.txt")

	story := CreateStory(filepaths, false, false)
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

func ExampleStory_common() {
	filepaths := []string{}
	filepaths = append(filepaths, paths...)
	filepaths = append(filepaths, "testdata/common.txt")

	story := CreateStory(filepaths, true, true)
	story.SetWriter(os.Stdout)
	story.Test()
	// Output:
	// Creating Room as kind of object.
	// Creating Thing as kind of object.
	// Creating Door as kind of thing.
	// Creating Container as kind of thing.
	// Creating Vehicle as kind of container.
	// Creating Supporter as kind of thing.
	// Creating Backdrop as kind of object.
	// Creating Person as kind of thing.
	// Creating Man as kind of person.
	// Creating Woman as kind of person.
	// Creating Animal as kind of person.
	// Creating Device as kind of thing.
	// Creating Direction as kind of object.
	// Creating Region as kind of object.
	// Creating dead end as kind of room.

}
