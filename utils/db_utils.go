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

func (u DBUtils) Read(db *sql.DB, query string, arguments ...interface{}) (*sql.Rows, error) {
	fmt.Println("== Reading Data From Database ==")
	result, err := db.Query(query, arguments...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func (u DBUtils) Insert(db *sql.DB, query string, arguments ...interface{}) (sql.Result, error) {
	fmt.Println("== Insert Data From Database ==")
	insert, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	result, err := insert.Exec(arguments...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func (u DBUtils) Update(db *sql.DB, query string, arguments ...interface{}) (sql.Result, error) {
	fmt.Println("== Update Data From Database ==")
	update, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	result, err := update.Exec(arguments...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u DBUtils) Delete(db *sql.DB, query string, arguments ...interface{}) (sql.Result, error) {
	fmt.Println("== Delete Data From Database ==")
	update, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	result, err := update.Exec(arguments...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
