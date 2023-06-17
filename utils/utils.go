package utils

import "github.com/spf13/viper"

type Config struct {
	Workspace_path        string   `mapstructure:"WORKSPACE_PATH"`
	Cme_db_path           string   `mapstructure:"CME_DB_PATH"`
	Variables_prefix      string   `mapstructure:"VARIABLES_PREFIX"`
	Variables_custom_list []string `mapstructure:"VARIABLES_CUSTOM_LIST"`
	Aliases_file_path     string   `mapstructure:"ALIASES_FILE_PATH"`
	Local_database_name   string   `mapstructure:"LOCAL_DATABASE_NAME"`
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
