package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"project_template/logger"
	"project_template/utils"
)

var database *gorm.DB

func initDB() *gorm.DB {
	settingParams := utils.AppSettings.PostgresParams

	connString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		settingParams.Server,
		settingParams.Port,
		settingParams.User,
		settingParams.DataBase,
		settingParams.Password)

	db, err := gorm.Open("postgres", connString)
	if err != nil {
		logger.Error.Fatal("Couldn't connect to database", err.Error())
	}

	// enabling gorm log mode, used for debugging
	db.LogMode(true)

	db.SingularTable(true)

	return db
}

// StartDbConnection Creates connection to database
func StartDbConnection() {
	database = initDB()
}

// GetDBConn func for getting db conn globally
func GetDBConn() *gorm.DB {
	return database
}
