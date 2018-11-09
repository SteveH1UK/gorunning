package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SteveH1UK/gorunning"
	"github.com/SteveH1UK/gorunning/validation"
)

func (env *Env) createAthelete(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Creating athelete")

	athelete := root.NewAthelete{}

	err := callDecodeJSONFromBody(r.Body, &athelete)
	if err != nil {
		log.Println("Error in decoding althelete from json")
		errorWithJSON(w, "Can not decode request", http.StatusBadRequest, nil)
		return
	}

	fmt.Println("Athelete to be created", athelete)

	validationErrors := validation.ValidateAthelete(athelete)
	if len(validationErrors) > 0 {
		errorWithJSON(w, "Validation Errors", http.StatusUnprocessableEntity, validationErrors)
		return
	}

	log.Printf("New Athelete Struct %+v", athelete)
	err = env.atheleteDAO.CreateAthelete(&athelete)
	if err != nil {
		log.Println("ERROR creating Athelete ", athelete, " err:", err)
		if err == root.ErrDBRecordExists {
			errorWithJSON(w, "Athelete with this friendly names already exists", http.StatusUnprocessableEntity, nil)
		} else {
			errorWithJSON(w, "Database error", http.StatusInternalServerError, nil)
		}
	} else {
		log.Println("Created Athelete ", athelete)
		location := r.URL.String() + "/" + athelete.FriendyURL
		newResource := root.NewResource{ID: athelete.FriendyURL, HTTPCode: http.StatusCreated, Href: location}
		respBody, _ := json.Marshal(newResource)
		responseWithJSON(w, respBody, http.StatusCreated)
	}
}
