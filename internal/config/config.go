package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	GrpcServer    string `mapstructure:"GRPC_SERVER"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	DB            *gorm.DB
	JwtExpiration int `mapstructure:"JWT_EXPIRATION"`
}

func LoadConfig() (Config, error) {
	var cfg Config

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return cfg, err
	}

	err = viper.Unmarshal(&cfg)
	return cfg, err
}
