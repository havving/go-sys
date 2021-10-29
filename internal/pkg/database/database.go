package database

import (
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	// postgres driver
	_ "github.com/lib/pq"
)

const (
	HOST     = ""
	DATABASE = ""
	USER     = ""
	PASSWORD = ""
)

func Setup() {
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", HOST, USER, PASSWORD, DATABASE)

	db, err := sql.Open("postgres", conn)
	defer db.Close()

	if err != nil || db.Ping() != nil {
		panic(err.Error())
	}
	log.Info().Msg("Success connection to DB")

}
