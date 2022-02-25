package main

import (
	"project_template/db"
	"project_template/logger"
	"project_template/routes"
	"project_template/utils"
)

func main() {
	utils.ReadSettings()
	utils.PutAdditionalSettings()
	logger.Init()
	db.StartDbConnection()
	//go jobs.StartJobs()
	routes.RunAllRoutes()
}
