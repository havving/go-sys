package database

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"

	// postgres driver
	_ "github.com/lib/pq"
	"log"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	DATABASE = "sys_mgr"
	USER     = "heybin"
	PASSWORD = "1234"
	SSL      = "disable"
)

type SQL struct {
	db   *sqlx.DB
	Exec func(query string, args ...interface{}) (sql.Result, error)
}

type SQLTX struct {
	tx       *sqlx.Tx
	Exec     func(query string, args ...interface{}) (sql.Result, error)
	Commit   func() error
	Rollback func() error
}

func Setup() {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", HOST, PORT, USER, PASSWORD, DATABASE, SSL)

	db, err := sql.Open("postgres", conn)
	defer db.Close()

	if err != nil || db.Ping() != nil {
		panic(err.Error())
	}
	log.Println("Success connection to DB")
}
