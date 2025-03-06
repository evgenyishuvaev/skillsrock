package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig
	DB  DBConfig
}

type AppConfig struct {
	Port string `mapstructure:"APP_PORT"`
	// MigrationsDir string
}

func (a *AppConfig) AsFiberPort() string {
	return fmt.Sprintf(":%s", a.Port)
}

type DBConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	DBName   string `mapstructure:"DB_NAME"`
}

func (d *DBConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", d.User, d.Password, d.Host, d.Port, d.DBName)
}

func NewConfig() *Config {

	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	appConfig := AppConfig{}
	dbConfig := DBConfig{}

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	return &Config{App: appConfig, DB: dbConfig}
}
