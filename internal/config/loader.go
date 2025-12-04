package config

import (
	"log"
	"time"
	"github.com/spf13/viper"
)

var GlobalConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	
	viper.SetConfigType("yaml")
	
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error reading config file: %v", err)
	}

	port := viper.GetInt("server.port")
	host := viper.GetString("server.host")

	fmt.Printf("server: %s:%d\n", host, port)

}