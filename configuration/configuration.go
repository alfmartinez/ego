package configuration

type Mob struct {
	Name      string   `yaml:"name"`
	Age       int      `yaml:"age"`
	Templates []string `yaml:"templates"`
}

type Template struct {
	Name       string   `yaml:"name"`
	Behaviours []string `yaml:"behaviours"`
}

type Configuration struct {
	ModelVersion string `yaml:"modelVersion"`
	Templates    []Template
	Mobs         []Mob
}
