package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SteveH1UK/gorunning"

	"github.com/SteveH1UK/gorunning/mappers"
)

func (env *Env) getAthelete(w http.ResponseWriter, r *http.Request) {

	friendlyName := getFriendlyName(r)
	log.Println("Friendly name", friendlyName)

	atheleteModel, err := env.atheleteDAO.FindAtheleteByFriendlyName(friendlyName)
	if err == nil {
		athelete := mappers.NewAtheleteFromModel(atheleteModel, "/gorunning/athelete/")
		log.Printf("Find athelete [%v] \n", athelete)
		log.Printf("URL schema [%+s] host[%s] path[%s] \n", r.URL.Scheme, r.URL.Host, r.URL.Path)

		respBody, _ := json.Marshal(athelete)
		responseWithJSON(w, respBody, http.StatusOK)
	} else {
		switch err {
		default:
			log.Println("DB Error: ", err)
			errorWithJSON(w, "Database error", http.StatusInternalServerError, nil)
		case root.ErrDBErrNotFound:
			log.Println("Error: ", err)
			errorWithJSON(w, "Athelete not found", http.StatusNotFound, nil)
		}
	}
}
