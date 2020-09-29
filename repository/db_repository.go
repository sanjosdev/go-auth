package repository

import (
	. "sanjos/auth/interfaces"
	"sanjos/auth/utils"
)

func DBRepository() DBInterface {

	// TODO : getting environment value here
	dbutils := utils.DBUtils{
		DBName:   "realcast-auth",
		Driver:   "mysql",
		Password: "",
		Username: "root",
	}

	return &dbutils

}
