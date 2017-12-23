package bikescrape

import (
	"os"
)

var RavenContext map[string]string = make(map[string]string)
var Username string
var Password string
var Host string
var DBName string

func init() {
	env, envExists := os.LookupEnv("ENVIRONMENT")
	Username = os.Getenv("POSTGRES_USER")
	Password = os.Getenv("POSTGRES_PASSWORD")
	Host = os.Getenv("POSTGRES_HOST")
	DBName = os.Getenv("POSTGRES_DB_NAME")

	if !envExists {
		env = "development"
	}

	RavenContext["environment"] = env
}
