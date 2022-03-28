package configuration

import (
	"github.com/alfmartinez/ego/engine/context"

	"github.com/spf13/viper"
)

func FromContext() *viper.Viper {
	return context.GetContext().Get("cfg").(*viper.Viper)
}

type ConfigurationFactory interface {
	Init()
	Get() *viper.Viper
}

func CreateConfiguration(path string) ConfigurationFactory {
	v := viper.New()
	v.BindEnv("lang")
	return &configuration{v, path}
}

type configuration struct {
	viper *viper.Viper
	path  string
}

func (c *configuration) Get() *viper.Viper {
	return c.viper
}

func (c *configuration) Init() {
	err := c.readConfig()
	if err != nil {
		panic(err)
	}
}

func (c *configuration) readConfig() error {
	err := c.importConfig([]string{"config"})
	if err != nil {
		return err
	}
	return nil
}

func (c *configuration) importConfig(imports []string) error {
	for _, x := range imports {
		v := viper.New()
		v.SetConfigName(x)
		v.AddConfigPath(c.path)
		err := v.ReadInConfig()
		if err != nil {
			return err
		}

		c.viper.MergeConfigMap(v.AllSettings())

		imports := v.GetStringSlice("imports")
		err = c.importConfig(imports)
		if err != nil {
			return err
		}
	}
	return nil
}
