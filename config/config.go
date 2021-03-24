package config

import (
	"path"

	"github.com/open-cmi/goutils/common"
	"github.com/spf13/viper"
)

// Conf conf
var Conf *viper.Viper

type ConfigOption struct {
	ConfigPath string
	FileName   string
	Format     string
}

var option ConfigOption

func SetConfigOption(configPath string, filename string, format string) {
	option.ConfigPath = configPath
	option.FileName = filename
	option.Format = format
	return
}

// InitConfig init config
func InitConfig() (err error) {
	Conf = viper.New()
	// init default option
	if option.ConfigPath == "" {
		root := common.GetRootPath()
		option.ConfigPath = path.Join(root, "etc")
	}
	if option.FileName == "" {
		option.FileName = "config"
	}
	if option.Format == "" {
		option.Format = "yaml"
	}

	Conf.AddConfigPath(option.ConfigPath)
	Conf.SetConfigName(option.FileName)
	Conf.SetConfigType(option.Format)
	if err = Conf.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
