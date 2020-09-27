package controller

import (
	"database/sql"
	"log"
	. "sanjos/auth/interfaces"
	"sanjos/auth/repository"

	"github.com/google/uuid"
)

var dbInterface DBInterface
var dbInstance *sql.DB

func InitRepository() {
	dbInterface = repository.DBRepository()
	dbInstance = dbInterface.Connect()
}

func GetUserData() {
	//result, err := dbInterface.Read(dbInstance, "");
}

func InsertUserData() {
	id := uuid.New()
	_, err := dbInterface.Insert(dbInstance, "INSERT INTO users VALUES(?, ?, ?, ?, ?)", id, "ilzammulkhaq85@gmail.com", "suku", "ini_url", "jkjk")
	if err != nil {
		log.Fatal()
	}
}

func UpdateUserData() {

}

func DeleteUserData() {

}
