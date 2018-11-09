package http

import (
	"log"
	"net/http"

	"github.com/SteveH1UK/gorunning/config"
	"github.com/gorilla/pat"
	"github.com/justinas/alice"
)

const urlPrefix = "/gorunning"

// ServeHTTP facade for application HTTP Server
func ServeHTTP(env *Env) {

	router := pat.New()
	router.Post(urlPrefix+"/atheletes", env.createAthelete)
	router.Get(urlPrefix+"/athelete/{friendly_name}", env.getAthelete)
	router.Get(urlPrefix+"/atheletes", env.getAtheletes)
	router.Put(urlPrefix+"/athelete/{friendly_name}", env.editAthelete)

	cfg, err := config.Get()
	if err != nil {
		log.Println("Error getting config:", err)
		panic(1)
	}
	listenAddress := cfg.ListenAddress
	log.Println("Listen Address [", listenAddress, "]")

	chain := alice.New(handlerLogger)
	err = http.ListenAndServe(listenAddress, chain.Then(router))
	if err != nil {
		log.Print("Fail to start HTTP Server: ", err)
		panic(-1)
	}
}
