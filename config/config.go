package config

import (
	"log"
	"os"

	"github.com/ian-kent/gofigure"
)

type config struct {
	gofigure      interface{} `order:"env,flag"`
	DbName        string      `env:"RUNNING_DB_NAME" flag:"db-name" flagDesc:"Database Name"`
	MongoURL      string      `env:"RUNNING_MONGO_URL" flag:"mongo-url" flagDesc:"MongoDB URL"`
	ListenAddress string      `env:"RUNNING_LISTEN_ADDRESS" flag:"listen-address" flagDesc:"Application Listen Address"`
	TestDbName    string      `env:"TEST_RUNNING_DB_NAME" flag:"test-db-name" flagDesc:"Test Database Name"`
	TestMongoURL  string      `env:"TEST_RUNNING_MONGO_URL" flag:"test-mongo-url" flagDesc:"Test MongoDB URL"`
}

var cfg *config

// Get configures the application and returns the configuration
func Get() (*config, error) {

	if cfg != nil {
		return cfg, nil
	}

	log.Print("Creating a new configuration variable")
	cfg = &config{
		DbName:        "gorunning",
		MongoURL:      "localhost:27017",
		ListenAddress: "127.0.0.1:8080",
		TestDbName:    "testgorunning",
		TestMongoURL:  "localhost:27017",
	}
	// Enable multiple values in environment variables...
	err := os.Setenv("GOFIGURE_ENV_ARRAY", "1")
	if err != nil {
		// Not that Setenv can fail....
		return nil, err
	}
	err = gofigure.Gofigure(cfg)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return cfg, nil
}
