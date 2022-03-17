package configuration

import "github.com/spf13/viper"

type Configuration interface {
	Init()
	Get() *viper.Viper
}

func CreateConfiguration() Configuration {
	v := viper.New()
	return &configuration{v}
}

type configuration struct {
	viper *viper.Viper
}

func (c *configuration) Get() *viper.Viper {
	return c.viper
}

func (c *configuration) Init() {
	c.readConfig()
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
		v.AddConfigPath("assets/config/")
		v.AddConfigPath("../../assets/config/")
		err := v.ReadInConfig()
		if err != nil {
			return err
		}

		err = c.viper.MergeConfigMap(v.AllSettings())
		if err != nil {
			return err
		}

		imports := v.GetStringSlice("imports")
		err = c.importConfig(imports)
		if err != nil {
			return err
		}
	}
	return nil
}
