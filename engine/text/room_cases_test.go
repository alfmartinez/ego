package text

var roomCases = []ParserCase{
	{
		"Title",
		args{`"Just a Title"
`},
		World{
			[]*Statement{
				{
					Title: "Just a Title",
				},
			},
		},
	},
	{
		"Comment",
		args{`// My Comment
"Commenting Park"
`},
		World{
			[]*Statement{
				{
					Title: "Commenting Park",
				},
			},
		},
	},
	{
		"A Simple Room",
		args{`"A Room With A View"

Chambord est un lieu.
`},
		World{
			[]*Statement{
				{
					Title: "A Room With A View",
				},
				{
					Room: &Room{
						Designator: &Designator{[]string{"Chambord"}},
					},
				},
			},
		},
	},
	{
		"A Simple Room - multiple token designator",
		args{`"A Room With A View"
Château Chambord est un lieu.
`},
		World{
			[]*Statement{
				{
					Title: "A Room With A View",
				},
				{
					Room: &Room{
						Designator: &Designator{[]string{"Château", "Chambord"}},
					},
				},
			},
		},
	},
	{
		"A Room with complex designator and a description",
		args{`"A Room With A Nice View"
Le Château de Chambord est un lieu. "L'entrée du Château est majestueuse."
`},
		World{
			[]*Statement{
				{
					Title: "A Room With A Nice View",
				},
				{
					Room: &Room{
						Designator:  &Designator{[]string{"Le", "Château", "de", "Chambord"}},
						Description: "L'entrée du Château est majestueuse.",
					},
				},
			},
		},
	},
	{
		"Let's add an item",
		args{`"A Room With A View"
Le Château de Chambord est un lieu. "L'entrée du Château est majestueuse."
Un livre déchiré est ici.
`},
		World{
			[]*Statement{
				{
					Title: "A Room With A View",
				},
				{
					Room: &Room{
						Designator:  &Designator{[]string{"Le", "Château", "de", "Chambord"}},
						Description: "L'entrée du Château est majestueuse.",
					},
				},
				{
					Item: &Item{
						Designator: &Designator{[]string{"Un", "livre", "déchiré"}},
					},
				},
			},
		},
	},
	{
		"Let's add an item - with description",
		args{`"A Room With A View"
Le Château de Chambord est un lieu. "L'entrée du Château est majestueuse."
Un livre déchiré est ici. "Un livre déchiré semble avoir été abandonné par un lecteur peu soigneux."
`},
		World{
			[]*Statement{
				{
					Title: "A Room With A View",
				},
				{
					Room: &Room{
						Designator:  &Designator{[]string{"Le", "Château", "de", "Chambord"}},
						Description: "L'entrée du Château est majestueuse.",
					},
				},
				{
					Item: &Item{
						Designator:  &Designator{[]string{"Un", "livre", "déchiré"}},
						Description: "Un livre déchiré semble avoir été abandonné par un lecteur peu soigneux.",
					},
				},
			},
		},
	},
	{
		"Add description after an item is created",
		args{`"A Room, a Table"
La cuisine est un lieu.
Une table est ici.
La description de la table est "Une vieille table en bois ancien." 				
`},
		World{
			[]*Statement{
				{
					Title: "A Room, a Table",
				},
				{
					Room: &Room{
						Designator: &Designator{[]string{"La", "cuisine"}},
					},
				},
				{
					Item: &Item{
						Designator: &Designator{[]string{"Une", "table"}},
					},
				},
				{
					Describe: &Describe{
						Designator:  &Designator{[]string{"de", "la", "table"}},
						Description: "Une vieille table en bois ancien.",
					},
				},
			},
		},
	},
}
