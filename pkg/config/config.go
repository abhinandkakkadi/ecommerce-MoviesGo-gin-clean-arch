package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	AUTHTOKEN   string `mapstructure:"TWILIO_AUTHTOKEN"`
	ACCOUNTSID  string `mapstructure:"TWILIO_ACCOUNTSID"`
	SERVICESSID string `mapstructure:"TWILIO_SERVICESID"`

	AWSS3REGION     string `mapstructure:"AWSS3_REGION"`
	AWSS3ACCESSKEY  string `mapstructure:"AWSS3_ACCESSKEY"`
	AWSS3SECRECTKEY string `mapstructure:"AWSS3_SECRECTKEY"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "TWILIO_AUTHTOKEN", "TWILIO_ACCOUNTSID", "TWILIO_SERVICESSID", "AWSS3_REGION", "AWSS3_ACCESSKEY", "AWSS3_SECRECTKEY",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil

}
