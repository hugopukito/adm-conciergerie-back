package repository

import (
	"adame/entity"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var DB *sql.DB
var err error

func InitDB() {
	DB, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/")
	if err != nil {
		log.Panicln("Failed to connect to sql: " + err.Error())
	}

	dbCreation := "CREATE DATABASE IF NOT EXISTS `adm-conciergerie`;"
	tableCreation := "CREATE TABLE IF NOT EXISTS form (ID BINARY(16) PRIMARY KEY, date TIMESTAMP, property_type VARCHAR(10), surface INT, property_condition VARCHAR(10), mail VARCHAR(255), phone VARCHAR(20), notes TEXT);"

	_, err = DB.Exec(dbCreation)
	if err != nil {
		log.Panicln("Failed to execute SQL statement: " + err.Error())
	}

	DB, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/adm-conciergerie?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	_, err = DB.Exec(tableCreation)
	if err != nil {
		log.Panicln("Failed to execute SQL statement: " + err.Error())
	}
}

func FindForms() ([]entity.Form, error) {
	selectQuery := "SELECT bin_to_uuid(id), date, property_type, surface, property_condition, mail, phone, notes FROM form ORDER BY date DESC"

	results, err := DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}

	var forms []entity.Form

	for results.Next() {
		var form entity.Form

		err = results.Scan(&form.ID, &form.Date, &form.PropertyType, &form.Surface, &form.PropertyCondition, &form.Mail, &form.Phone, &form.Notes)
		if err != nil {
			return nil, err
		}

		forms = append(forms, form)
	}

	return forms, nil
}

func InsertForm(form entity.Form) (uuid.UUID, error) {
	stmt, err := DB.Prepare("INSERT INTO form (id, date, property_type, surface, property_condition, mail, phone, notes) values (uuid_to_bin(?), ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return uuid.UUID{}, err
	}

	id := uuid.New()
	date := time.Now()

	_, err = stmt.Exec(id.String(), date, form.PropertyType, form.Surface, form.PropertyCondition, form.Mail, form.Phone, form.Notes)

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}
