package text

func ExampleStory_Simple() {
	filepath := "testdata/simple.txt"
	story := CreateStory(filepath)
	story.Test()

	// Output:
	// The Black Room
	//
}
