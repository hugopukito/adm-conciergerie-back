package service

import (
	"adame/cors"
	"adame/entity"
	"adame/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"

	"github.com/google/uuid"
)

var mailPwd string

func init() {
	filePath := "pwd.txt"

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	mailPwd = string(content)
}

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
		errMsg := fmt.Sprintf("invalid PropertyType || PropertyCondition || Surface: %s", form.PropertyType)
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
	go sendMail(form)
}

func sendMail(form entity.Form) {
	from := "senderadame3@gmail.com"
	to := "gogo26hugop@gmail.com"
	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	subject := "New Property Form Submission"
	body := fmt.Sprintf("Type de bien: %s\nSurface: %d sq. meters\nEtat du bien: %s\n"+
		"Email: %s\nPhone: %s\nNotes: %s\n",
		form.PropertyType, form.Surface, form.PropertyCondition,
		form.Mail, form.Phone, form.Notes)
	message := []byte("Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	auth := smtp.PlainAuth("", from, mailPwd, smtpServer)
	err := smtp.SendMail(smtpServer+":"+strconv.Itoa(smtpPort), auth, from, []string{to}, message)
	if err != nil {
		fmt.Println(err)
		return
	}
}
