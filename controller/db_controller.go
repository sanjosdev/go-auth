package controller

import (
	"log"
	. "sanjos/auth/interfaces"
	"sanjos/auth/repository"

	"github.com/google/uuid"
)

var dbInterface DBInterface

func InitRepository() {
	dbInterface = repository.DBRepository()
	dbInterface.Connect()
}

func GetUserData() {
	//result, err := dbInterface.Read(dbInstance, "");
}

func InsertUserData() {
	id := uuid.New()
	_, err := dbInterface.Insert("INSERT INTO users VALUES(?, ?, ?, ?, ?)", id, "ilzammulkhaq65@gmail.com", "suku", "ini_url", "jkjk")
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateUserData() {

}

func DeleteUserData() {

}
