package text

var informerSentences = []ParserCase{
	{
		"Simple sentence",
		"The chamber is a room.\n",
		[]*Statement{
			{
				Sentence: &Sentence{
					DP: &DescriptionPhrase{Simple: &Noun{"chamber"}},
					VP: &VerbPhrase{
						"is",
						&DescriptionPhrase{Simple: &Noun{"room"}},
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
					DP: &DescriptionPhrase{Simple: &Noun{"chamber"}},
					VP: &VerbPhrase{
						"is",
						&DescriptionPhrase{Simple: &Noun{"room"}},
					},
				},
			},
			{
				Sentence: &Sentence{
					DP: &DescriptionPhrase{Simple: &Noun{"cat"}},
					VP: &VerbPhrase{
						"is",
						&DescriptionPhrase{Simple: &Noun{"here"}},
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
					DP: &DescriptionPhrase{Simple: &Noun{"cart"}},
					VP: &VerbPhrase{
						"is",
						&DescriptionPhrase{Complex: &ComplexPhrase{
							Simple: &Noun{"vehicle"},
							VP: &VerbPhrase{
								"is",
								&DescriptionPhrase{Simple: &Noun{"movable"}},
							},
						}},
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
					DP: &DescriptionPhrase{
						Simple: &Noun{"Janitor"},
					},
					VP: &VerbPhrase{
						"is",
						&DescriptionPhrase{
							Complex: &ComplexPhrase{
								Simple: &Noun{"person"},
								VP: &VerbPhrase{
									"has",
									&DescriptionPhrase{
										AdjNoun: &AdjNoun{
											Adjective: "blue",
											Noun:      &Noun{"key"},
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
						Simple: &Noun{"cat"},
					},
					VP: &VerbPhrase{
						"is",
						&DescriptionPhrase{
							Relation: &RelativePhrase{
								Simple:   &Noun{"kind"},
								Relation: "of",
								Related:  &Noun{"animal"},
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
						AdjNoun: &AdjNoun{
							Adjective: "Blue",
							Noun:      &Noun{"Fury"},
						},
					},
					VP: &VerbPhrase{
						"is",
						&DescriptionPhrase{
							Simple: &Noun{"cat"},
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
						AdjNoun: &AdjNoun{
							Adjective: "Blue",
							Noun:      &Noun{"Fury"},
						},
					},
					VP: &VerbPhrase{
						"is",
						&DescriptionPhrase{
							Simple: &Noun{"cat"},
						},
					},
					Description: "A big cat lies lazily on the couch.",
				},
			},
		},
		false,
	},
}
