package interfaces

import "database/sql"

type DBInterface interface {
	Connect() *sql.DB
	Read(db *sql.DB, query string, arguments ...interface{}) (*sql.Rows, error)
	Insert(db *sql.DB, query string, arguments ...interface{}) (sql.Result, error)
	Update(db *sql.DB, query string, arguments ...interface{}) (sql.Result, error)
	Delete(db *sql.DB, query string, arguments ...interface{}) (sql.Result, error)
}
