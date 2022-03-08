package configuration

type Position struct {
	X int `yaml:"x"`
	Y int `yaml:"y"`
}

type Size struct {
	Width  int `yaml:"width"`
	Height int `yaml:"height"`
}

type Mob struct {
	Type     string   `yaml:"type"`
	Name     string   `yaml:"name"`
	Position Position `yaml:"position"`
	Sprite   string   `yaml:"sprite"`
	Needs    []Need
}

type Renderer struct {
	Type    string  `yaml:"type"`
	Display Display `yaml:"display"`
}

type Display struct {
	Type     string `yaml:"type"`
	Size     Size   `yaml:"size"`
	ViewPort Size   `yaml:"viewport"`
}

type Need struct {
	Type  string `yaml:"name"`
	Level int    `yaml:"level"`
}

type Configuration struct {
	ModelVersion string   `yaml:"modelVersion"`
	Grid         Position `yaml:"grid"`
	Mobs         []Mob
	Renderer     Renderer `yaml:"renderer"`
}
