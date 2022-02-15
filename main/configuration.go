package main

type MobConfig struct {
	Name      string   `yaml:"name"`
	Age       int      `yaml:"age"`
	Templates []string `yaml:"templates"`
}

type TemplateConfig struct {
	Name       string   `yaml:"name"`
	Behaviours []string `yaml:"behaviours"`
}

type Configuration struct {
	ModelVersion string `yaml:"modelVersion"`
	Templates    []TemplateConfig
	Mobs         []MobConfig
}
