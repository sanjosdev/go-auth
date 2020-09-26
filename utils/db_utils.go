package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DBUtils struct {
	Driver   string
	Username string
	Password string
	DBName   string
}

func (u DBUtils) Connect() *sql.DB {
	fmt.Println("== Connecting To Database ==")
	dburl := fmt.Sprintf("%s:%s@/%s", u.Username, u.Password, u.DBName)
	db, err := sql.Open(u.Driver, dburl)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
