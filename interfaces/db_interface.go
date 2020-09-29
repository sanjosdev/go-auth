package interfaces

import "database/sql"

type DBInterface interface {
	Connect()
	Read(query string, arguments ...interface{}) (*sql.Rows, error)
	Insert(query string, arguments ...interface{}) (sql.Result, error)
	Update(query string, arguments ...interface{}) (sql.Result, error)
	Delete(query string, arguments ...interface{}) (sql.Result, error)
}
