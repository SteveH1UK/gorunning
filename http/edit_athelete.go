package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SteveH1UK/gorunning"
	"github.com/SteveH1UK/gorunning/mappers"
	"github.com/SteveH1UK/gorunning/validation"
)

func (env *Env) editAthelete(w http.ResponseWriter, r *http.Request) {

	friendlyName := getFriendlyName(r)
	fmt.Println("Editing athelete ", friendlyName)

	athelete := root.NewAthelete{}

	err := callDecodeJSONFromBody(r.Body, &athelete)
	if err != nil {
		log.Println("Error in decoding althelete from json")
		errorWithJSON(w, "Can not decode request", http.StatusBadRequest, nil)
		return
	}

	validationErrors := validation.ValidateAthelete(athelete)
	if len(validationErrors) > 0 {
		errorWithJSON(w, "Validation Errors", http.StatusUnprocessableEntity, validationErrors)
		return
	}

	log.Printf("Athelete Struct %+v", athelete)
	err = env.atheleteDAO.EditAthelete(friendlyName, &athelete)
	if err != nil {
		log.Println("ERROR creating Athelete ", athelete, " err:", err)
		if err == root.ErrDBRecordExists {
			errorWithJSON(w, "Athelete with updated friendly names already exists", http.StatusUnprocessableEntity, nil)
		} else {
			errorWithJSON(w, "Database error", http.StatusInternalServerError, nil)
		}
	} else {
		log.Println("Edit Athelete ", athelete)
		atheleteModel, err := env.atheleteDAO.FindAtheleteByFriendlyName(athelete.FriendyURL)
		if err == nil {
			athelete := mappers.NewAtheleteFromModel(atheleteModel, "/atheletes/")
			log.Printf("Edited athelete [%v] \n", athelete)
			log.Printf("URL schema [%+s] host[%s] path[%s] \n", r.URL.Scheme, r.URL.Host, r.URL.Path)
			respBody, _ := json.Marshal(athelete)
			responseWithJSON(w, respBody, http.StatusOK)
		} else {
			errorWithJSON(w, "General error", http.StatusInternalServerError, nil)
		}
	}
}
