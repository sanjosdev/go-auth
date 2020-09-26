package interfaces

import "database/sql"

type DBInterface interface {
	Connect() *sql.DB
}
