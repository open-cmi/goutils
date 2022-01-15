package config

import (
	"path"

	"github.com/open-cmi/goutils/pathutil"
	"github.com/spf13/viper"
)

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
}

// InitConfig init config
func InitConfig() (conf *viper.Viper, err error) {
	parser := viper.New()
	// init default option
	if option.ConfigPath == "" {
		root := pathutil.GetRootPath()
		option.ConfigPath = path.Join(root, "etc")
	}
	if option.FileName == "" {
		option.FileName = "config"
	}

	if option.Format == "" {
		option.Format = "yaml"
	}

	parser.AddConfigPath(option.ConfigPath)
	parser.SetConfigName(option.FileName)
	parser.SetConfigType(option.Format)
	if err = parser.ReadInConfig(); err != nil {
		return nil, err
	}

	return parser, nil
}
