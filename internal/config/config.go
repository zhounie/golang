package config

import (
	"time"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"server" yaml:"server"`
	Database DatabaseConfig `mapstructure:"database" yaml:"database"`
	Redis RedisConfig `mapstructure:"redis" yaml:"redis"`
	JWT JWTConfig `mapstructure:"jwt" yaml:"jwt"`
	Log LogConfig `mapstructure:"log" yaml:"log"`
}

type ServerConfig struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port string `mapstructure:"port" yaml:"port"`
	Mode string `mapstructure:"mode" yaml:"mode"`
}

type DatabaseConfig struct {
	Driver string `mapstructure:"driver" yaml:"driver"`
	Host string `mapstructure:"host" yaml:"host"`
	Port string `mapstructure:"port" yaml:"port"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	Database string `mapstructure:"database" yaml:"database"`
	Charset string `mapstructure:"charset" yaml:"charset"`
}

type RedisConfig struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port string `mapstructure:"port" yaml:"port"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret" yaml:"secret"`
	Expire time.Duration `mapstructure:"expire" yaml:"expire"`
}

type LogConfig struct {
	Level string `mapstructure:"level" yaml:"level"`
}

var (
	globalConfig *config
	viperInstance *viper.Viper
)

func Init(configPath ...string) (*Config, error) {
	viperInstance = viper.New()

	
}