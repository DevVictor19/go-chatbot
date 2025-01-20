package api

import (
	"server/src/api/database"
	"server/src/api/routes"
)

func InitHttpServer() {
	database.InitDatabaseConnection()
	routes.InitRoutes()
}
