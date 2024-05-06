package configs

import (
	"fmt"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBPort                 int
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		PublicHost:             "http://localhost",
		Port:                   "8080",
		DBUser:                 "root",
		DBPassword:             "mypassword",
		DBAddress:              fmt.Sprintf("%s:%s", "127.0.0.1", "3306"),
		DBPort:                 3306,
		DBName:                 "ecom",
		JWTSecret:              "not-so-secret-now-is-it?",
		JWTExpirationInSeconds: 3600 * 24 * 7,
	}
}
