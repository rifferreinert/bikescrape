package bikescrape

import (
	"os"

	"github.com/rifferreinert/bikescrape/pkg/divvy/models"
)

var RavenContext map[string]string = make(map[string]string)
var DbConfig models.DBConfig

func init() {
	env, envExists := os.LookupEnv("ENVIRONMENT")
	user, userExists := os.LookupEnv("POSTGRES_USER")
	password, passwordExists := os.LookupEnv("POSGTRES_PASSWORD")
	host, hostExists := os.LookupEnv("POSTGRES_HOST")
	dbname, dbnameExists := os.LookupEnv("POSTGRES_DB_NAME")

	if !envExists {
		env = "Development"
	}

	if userExists && passwordExists && hostExists && dbnameExists {
		DbConfig = models.DBConfig{
			Username: user,
			Password: password,
			Host:     host,
			DBName:   dbname,
		}
	}

	RavenContext["environment"] = env
}
