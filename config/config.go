package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var AppPort int

var DbRetries int

var PgConnStr string

var MockApi bool
var InfoApiUrl string

var errorsArray []error

func Load() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file not found")
	}

	AppPort = getEnvInt("APP_PORT")

	DbRetries = getEnvInt("DB_RETRIES")

	PgConnStr = fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v",
		getEnv("POSTGRES_USERNAME"),
		getEnv("POSTGRES_PASSWORD"),
		getEnv("POSTGRES_HOST"),
		getEnv("POSTGRES_PORT"),
		getEnv("POSTGRES_DATABASE"),
	)

	MockApi = getEnv("MOCK_API") == "YES"
	if MockApi {
		InfoApiUrl = fmt.Sprintf(
			"http://localhost:%v/info",
			AppPort,
		)
	} else {
		InfoApiUrl = getEnv("INFO_API_URL")
	}

	if len(errorsArray) > 0 {
		return errors.Join(errorsArray...)
	}

	return nil
}

func getEnv(envName string) string {
	val, ok := os.LookupEnv(envName)
	if !ok {
		errorsArray = append(errorsArray, fmt.Errorf("env %v was not found", envName))
		return ""
	}
	return val

}

func getEnvInt(envName string) int {
	str := getEnv(envName)
	if str == "" {
		errorsArray = append(errorsArray, fmt.Errorf("env %v was not found", envName))
		return 0
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		errorsArray = append(errorsArray, err)
		return 0
	}
	return val
}
