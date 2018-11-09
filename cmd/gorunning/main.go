package main

import (
	"log"

	"github.com/SteveH1UK/gorunning/config"
	"github.com/SteveH1UK/gorunning/http"
	"github.com/SteveH1UK/gorunning/mongodb"
)

func main() {

	log.Println("Starting application")

	cfg, err := config.Get()
	if err != nil {
		log.Println("Error getting config:", err)
		panic(1)
	}

	m, err := mongodb.Init(cfg.MongoURL, cfg.DbName)
	if err != nil {
		log.Fatalln("unable to connect to mongodb ", err)
		panic(-1)
	}
	defer m.Shutdown()

	atheleteDAO := mongodb.NewAtheleDAO(m)
	env := http.NewEnv(atheleteDAO)

	http.ServeHTTP(env)
}
