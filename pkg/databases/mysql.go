package databases

import (
	"database/sql"
	"time"

	"github.com/AasSuhendar/go-restful-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "admin:password@tcp(localhost:3306)/db")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
