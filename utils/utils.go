package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

const (
	DbUser     = "db_user"
	DbPassword = "db_pass"
	DbHost     = "db_host"
	DbName     = "db_name"
	Port       = "port"
)

func LoadEnvConfigs() (DbConnectionUri string) {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}
	username := os.Getenv(DbUser)
	password := os.Getenv(DbPassword)
	dbName := os.Getenv(DbName)
	dbHost := os.Getenv(DbHost)

	uri := username + ":" + password + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8&parseTime=True"
	return uri
}

func GetServicePort() (port string) {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}
	portNumber := os.Getenv(Port)
	return portNumber
}
