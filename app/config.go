package app

import (
	"os"

	"github.com/spf13/viper"
)

func (a *app) loadConfig() {
	filePath := os.Getenv("CONFIG_PATH")
	if filePath == "" {
		filePath = "./config.yaml"
	}
	reader, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	a.viper = viper.NewWithOptions(viper.KeyDelimiter("."))
	a.viper.SetConfigType("yaml")
	if err = a.viper.ReadConfig(reader); err != nil {
		panic(err)
	}
}
