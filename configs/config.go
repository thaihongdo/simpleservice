package configs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var cf *Configuration

func GetConfig() *Configuration {
	return cf
}

type Configuration struct {
	EnvPrefix        string `mapstructure:"env_prefix"`
	ServerPort       int    `mapstructure:"server_port"`
	DBConnection     string `mapstructure:"db_connection"`
	TokenSecret      string `mapstructure:"token_secret"`
	DefaultPageNum   int    `mapstructure:"default_page_num"`
	DefaultPageLimit int    `mapstructure:"default_page_limit"`
	LogConfig        LogConfig
}

type LogConfig struct {
	Format  string `mapstructure:"format"`
	Output  string `mapstructure:"output"`
	Expired int    `mapstructure:"expired"`
	Path    string `mapstructure:"path"`
}

//InitFromFile init config file
func InitFromFile(path string) *Configuration {
	if path == "" {
		viper.AddConfigPath("config")
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
	} else {
		viper.SetConfigFile(path)
	}

	//basePath, _ := os.Getwd()

	viper.AutomaticEnv()
	viper.SetEnvPrefix("")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		log.Printf("Config file not found: %v", err)
		panic(err)
	}
	//viper.Set("base_path", basePath)

	if err := viper.Unmarshal(&cf); err != nil {
		log.Printf("covert to struct: %v", err)
		log.Fatal(err)
	}
	if path == "" {
		fmt.Printf("File config used  %s\n \n", viper.ConfigFileUsed())
		dataPrinf, _ := json.Marshal(cf)
		fmt.Printf("Config:  %s\n \n", string(dataPrinf))
	}

	return cf
}
