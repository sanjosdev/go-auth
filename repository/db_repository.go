package repository

import (
	. "sanjos/auth/interfaces"
	"sanjos/auth/utils"
)

func DBRepository() DBInterface {

	dbutils := utils.DBUtils{
		DBName:   "",
		Driver:   "mysql",
		Password: "",
		Username: "root",
	}

	return dbutils

}
