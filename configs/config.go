package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
type db struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type Config struct {
	server server
	db     db
}

func (cfg *Config) GetConnectionString() string {
	db := cfg.db
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		db.Username, db.Password, db.Host, db.Port, db.DbName, db.SSLMode)
}
func (cfg *Config) GetAddress() string {
	return fmt.Sprintf(":%s", cfg.server.Port)
}
func New() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs/")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	server := server{
		Port: viper.GetString("server.port"),
	}
	db := db{
		DbName:   viper.GetString("db.dbname"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	return &Config{
		server: server,
		db:     db,
	}, nil
}
