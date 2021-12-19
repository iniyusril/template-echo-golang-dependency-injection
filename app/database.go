package app

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/iniyusril/template/helper"
	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	var (
		host     = os.Getenv("DB_POSTGREE_HOST")
		port     = os.Getenv("DB_POSTGREE_PORT")
		user     = os.Getenv("DB_POSTGREE_USER")
		password = os.Getenv("DB_POSTGREE_PASSWORD")
		dbname   = os.Getenv("DB_POSTGREE_DBNAME")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		helper.PanicIfError(err)
	}
	return db
}
