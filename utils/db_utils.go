package utils

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBUtils struct {
	Driver     string
	Username   string
	Password   string
	DBName     string
	DBInstance *sql.DB
}

func (u *DBUtils) Connect() {
	fmt.Println("== Connecting To Database ==")
	dburl := fmt.Sprintf("%s:%s@/%s", u.Username, u.Password, u.DBName)
	db, err := sql.Open(u.Driver, dburl)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	u.DBInstance = db
	fmt.Println(u.DBInstance)
}

func (u *DBUtils) Read(query string, arguments ...interface{}) (*sql.Rows, error) {
	fmt.Println("== Reading Data From Database ==")
	result, err := u.DBInstance.Query(query, arguments...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func (u *DBUtils) Insert(query string, arguments ...interface{}) (sql.Result, error) {
	fmt.Println("== Insert Data From Database ==")
	insert, err := u.DBInstance.Prepare(query)
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

func (u *DBUtils) Update(query string, arguments ...interface{}) (sql.Result, error) {
	fmt.Println("== Update Data From Database ==")
	update, err := u.DBInstance.Prepare(query)
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

func (u DBUtils) Delete(query string, arguments ...interface{}) (sql.Result, error) {
	fmt.Println("== Delete Data From Database ==")
	update, err := u.DBInstance.Prepare(query)
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
