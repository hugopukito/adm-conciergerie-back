package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type PropertyType string
type PropertyCondition string

const (
	House     PropertyType      = "house"
	Apartment PropertyType      = "apartment"
	Renovate  PropertyCondition = "renovate"
	Decorate  PropertyCondition = "decorate"
	Ready     PropertyCondition = "ready"
)

type Form struct {
	ID                uuid.UUID         `json:"ID,omitempty"`
	Date              time.Time         `json:"date,omitempty"`
	PropertyType      PropertyType      `json:"propertyType,omitempty"`
	Surface           int               `json:"surface,omitempty"`
	PropertyCondition PropertyCondition `json:"propertyCondition,omitempty"`
	Mail              string            `json:"mail,omitempty"`
	Phone             string            `json:"phone,omitempty"`
	Notes             string            `json:"notes,omitempty"`
}

func GetForms(w http.ResponseWriter, r *http.Request) {
	if EnableCors(&w, r) == "options" {
		return
	}

	var forms []Form
	var err error

	forms, err = FindForms()
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
	if EnableCors(&w, r) == "options" {
		return
	}

	var form Form
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if (form.PropertyType != House && form.PropertyType != Apartment) || (form.PropertyCondition != Renovate && form.PropertyCondition != Decorate && form.PropertyCondition != Ready || (form.Surface < 15)) {
		errMsg := fmt.Sprintf("invalid PropertyType value: %s", form.PropertyType)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(errMsg))
		return
	}

	var id uuid.UUID
	id, err = InsertForm(form)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", id.String())
	w.WriteHeader(http.StatusCreated)
}
