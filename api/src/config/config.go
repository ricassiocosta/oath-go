package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DBConnectionString = ""
	APIPort            = 0
	APISecret          []byte
	RootCACertPath     = ""
	RootCAKeyPath      = ""
	GithubClientID     = ""
	GithubClientSecret = ""
)

const (
	defaultAPIPort = 5000
)

// LoadEnv will initialize the environment vars
func LoadEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	APIPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		APIPort = defaultAPIPort
	}

	APISecret = []byte(os.Getenv("API_SECRET"))
	RootCACertPath = os.Getenv("ROOT_CA_CERT_PATH")
	RootCAKeyPath = os.Getenv("ROOT_CA_KEY_PATH")
	GithubClientID = os.Getenv("GITHUB_CLIENT_ID")
	GithubClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")

	dbAddress := os.Getenv("DB_ADDRESS")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	DBConnectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbAddress,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)
}
