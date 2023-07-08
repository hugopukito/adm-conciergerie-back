package service

import (
	"adame/cors"
	"adame/entity"
	"adame/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func GetForms(w http.ResponseWriter, r *http.Request) {
	if cors.EnableCors(&w, r) == "options" {
		return
	}

	var forms []entity.Form
	var err error

	forms, err = repository.FindForms()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	dto, err := json.Marshal(forms)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(string(dto)))
}

func PostForm(w http.ResponseWriter, r *http.Request) {
	if cors.EnableCors(&w, r) == "options" {
		return
	}

	var form entity.Form
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if (form.PropertyType != entity.House && form.PropertyType != entity.Apartment) || (form.PropertyCondition != entity.Renovate && form.PropertyCondition != entity.Decorate && form.PropertyCondition != entity.Ready || (form.Surface < 15)) {
		errMsg := fmt.Sprintf("invalid PropertyType value: %s", form.PropertyType)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(errMsg))
		return
	}

	var id uuid.UUID
	id, err = repository.InsertForm(form)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", id.String())
	w.WriteHeader(http.StatusCreated)
}
