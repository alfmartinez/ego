package text

var definitions = []ParserCase{
	{
		"Section - title with space",
		`Section 1 - "Testing descriptions"
`,
		[]*Statement{
			{Section: &Section{
				Number: 1,
				Title:  "Testing descriptions",
			}},
		},
		false,
	},
	{
		"Section - one word title",
		`Section 2 - Story
`,
		[]*Statement{
			{Section: &Section{
				Number: 2,
				Title:  "Story",
			}},
		},
		false,
	},
	{
		"Section - one word title",
		`Section 3 - Testing - Not For Release
`,
		[]*Statement{
			{Section: &Section{
				Number: 3,
				Title:  "Testing",
				Tag:    "NotForRelease",
			}},
		},
		false,
	},
}
