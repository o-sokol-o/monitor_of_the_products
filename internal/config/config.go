package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB     Mongo
	Server Server
}

type Mongo struct {
	URI      string `default:"mongodb://localhost:27017"`
	Username string `default:"admin"`
	Password string `default:"g0langn1nja"`
	Database string `default:"monitor_of_the_products"`
}

type Server struct {
	Port int `default:"9000"`
}

func New() (*Config, error) {

	// Инициализация читалки конфига
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("configs/") // path to look for the config file in
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		log.Fatalf("fatal error config file: %+v", err)
	}

	cfg := Config{
		DB: Mongo{
			URI:      viper.GetString("db.uri"),
			Username: viper.GetString("db.login"),
			Password: viper.GetString("db.password"),
			Database: viper.GetString("db.database"),
		},
		Server: Server{
			Port: viper.GetInt("Server.Port"),
		},
	}

	return &cfg, nil
}
