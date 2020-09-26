package controller

import (
	"database/sql"
	. "sanjos/auth/interfaces"
	"sanjos/auth/repository"
)

var dbInterface DBInterface
var dbInstance *sql.DB

func InitRepository() {
	dbInterface = repository.DBRepository()
	dbInstance = dbInterface.Connect()
}

func ReadData() {

}
