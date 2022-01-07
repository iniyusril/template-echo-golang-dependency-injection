package app

import (
	"fmt"
	"os"
	"strconv"

	"github.com/iniyusril/template/helper"
	"github.com/iniyusril/template/model/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func NewDB() *gorm.DB {
	var (
		host     = os.Getenv("DB_POSTGREE_HOST")
		port     = os.Getenv("DB_POSTGREE_PORT")
		user     = os.Getenv("DB_POSTGREE_USER")
		password = os.Getenv("DB_POSTGREE_PASSWORD")
		dbname   = os.Getenv("DB_POSTGREE_DBNAME")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)

	db, err := gorm.Open("postgres", dsn)

	isMigrate, err := strconv.ParseBool(os.Getenv("DB_POSTGREE_IS_AUTO_MIGRATE"))
	if isMigrate && err == nil {
		db.AutoMigrate(&domain.Category{})
	}

	if err != nil {
		helper.PanicIfError(err)
	}
	return db
}
