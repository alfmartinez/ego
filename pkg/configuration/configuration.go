package configuration

type Position struct {
	X int `yaml:"x"`
	Y int `yaml:"y"`
}

type Mob struct {
	Name     string   `yaml:"name"`
	Position Position `yaml:"position"`
}

type Configuration struct {
	ModelVersion string   `yaml:"modelVersion"`
	Grid         Position `yaml:"grid"`
	Mobs         []Mob
}
