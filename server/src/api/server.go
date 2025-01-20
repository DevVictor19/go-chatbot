package api

func InitHttpServer() {
	initDatabaseConnection()
	initRoutes()
}
