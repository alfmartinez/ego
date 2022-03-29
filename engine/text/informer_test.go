package text

var informerSentences = []ParserCase{
	{
		"Simple sentence",
		"The chamber is a room.\n",
		[]*Statement{
			{
				Sentence: &Sentence{
					DP: &DescriptionPhrase{Designator: &Designator{[]string{"chamber"}}},
					VP: &VerbPhrase{
						"is",
						&DescriptionPhrase{Designator: &Designator{[]string{"room"}}},
					},
				},
			},
		},
		false,
	},
	{
		"Several sentences",
		"The chamber is a room.\nA cat is here.\n",
		[]*Statement{
			{
				Sentence: &Sentence{
					DP: &DescriptionPhrase{Designator: &Designator{[]string{"chamber"}}},
					VP: &VerbPhrase{
						"is",
						&DescriptionPhrase{Designator: &Designator{[]string{"room"}}},
					},
				},
			},
			{
				Sentence: &Sentence{
					DP: &DescriptionPhrase{Designator: &Designator{[]string{"cat"}}},
					VP: &VerbPhrase{
						"is",
						&DescriptionPhrase{Designator: &Designator{[]string{"here"}}},
					},
				},
			},
		},
		false,
	},
	{
		"Sentence with determiner - which",
		"The cart is a vehicle which is movable.\n",
		[]*Statement{
			{
				Sentence: &Sentence{
					DP: &DescriptionPhrase{Designator: &Designator{[]string{"cart"}}},
					VP: &VerbPhrase{
						"is",
						&DescriptionPhrase{
							Designator: &Designator{[]string{"vehicle"}},
							Complex: &ComplexPhrase{
								VP: &VerbPhrase{
									"is",
									&DescriptionPhrase{Designator: &Designator{[]string{"movable"}}},
								},
							},
						},
					},
				},
			},
		},
		false,
	},
	{
		"Sentence with determiner - who",
		"The Janitor is a person who has the blue key.\n",
		[]*Statement{
			{
				Sentence: &Sentence{
					DP: &DescriptionPhrase{Designator: &Designator{[]string{"Janitor"}}},
					VP: &VerbPhrase{
						"is",
						&DescriptionPhrase{
							Designator: &Designator{[]string{"person"}},
							Complex: &ComplexPhrase{
								VP: &VerbPhrase{
									"has",
									&DescriptionPhrase{
										Designator: &Designator{
											[]string{"blue", "key"},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		false,
	},
	{
		"Kind sentence",
		"A cat is a kind of animal.\n",
		[]*Statement{
			{
				Sentence: &Sentence{
					DP: &DescriptionPhrase{
						Designator: &Designator{[]string{"cat"}},
					},
					VP: &VerbPhrase{
						"is",
						&DescriptionPhrase{
							Designator: &Designator{[]string{"kind"}},
							Relation: &RelativePhrase{
								Relation: "of",
								Related:  &Designator{[]string{"animal"}},
							},
						},
					},
				},
			},
		},
		false,
	},
	{
		"Adjective",
		"The Blue Fury is a cat.\n",
		[]*Statement{
			{
				Sentence: &Sentence{
					DP: &DescriptionPhrase{
						Designator: &Designator{[]string{"Blue", "Fury"}},
					},
					VP: &VerbPhrase{
						"is",
						&DescriptionPhrase{
							Designator: &Designator{[]string{"cat"}},
						},
					},
				},
			},
		},
		false,
	},
	{
		"Title",
		`"The Blue Fury is a cat."
`,
		[]*Statement{
			{
				Title: "The Blue Fury is a cat.",
			},
		},

		false,
	},
	{
		"Statement with optional description",
		`The Blue Fury is a cat. "A big cat lies lazily on the couch."
`,
		[]*Statement{
			{
				Sentence: &Sentence{
					DP: &DescriptionPhrase{
						Designator: &Designator{[]string{"Blue", "Fury"}},
					},
					VP: &VerbPhrase{
						"is",
						&DescriptionPhrase{
							Designator: &Designator{[]string{"cat"}},
						},
					},
					Description: "A big cat lies lazily on the couch.",
				},
			},
		},
		false,
	},
}
