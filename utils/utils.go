package utils

import "github.com/spf13/viper"

type Config struct {
	Workspace_path        string   `mapstructure:"WORKSPACE_PATH"`
	Cme_db_path           string   `mapstructure:"CME_DB_PATH"`
	Variables_prefix      string   `mapstructure:"VARIABLES_PREFIX"`
	Variables_custom_list []string `mapstructure:"VARIABLES_CUSTOM_LIST"`
	Exegol_aliases_path   string   `mapstructure:"EXEGOL_ALIASES_PATH"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
