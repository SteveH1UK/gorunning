package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SteveH1UK/gorunning"

	"github.com/SteveH1UK/gorunning/mappers"
)

func (env *Env) getAtheletes(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	log.Println("Params", params)

	friendlyName := params.Get(":friendly_name")
	log.Println("Friendly name", friendlyName)

	atheletesModel, err := env.atheleteDAO.FindAllAtheletes()
	if err == nil {
		var atheletes []root.Athelete
		for _, atheleteModel := range atheletesModel {
			athelete := mappers.NewAtheleteFromModel(atheleteModel, "/gorunning/athelete/")
			atheletes = append(atheletes, athelete)
		}
		respBody, _ := json.Marshal(atheletes)
		responseWithJSON(w, respBody, http.StatusOK)
	} else {
		errorWithJSON(w, "Database error", http.StatusInternalServerError, nil)
	}
}
