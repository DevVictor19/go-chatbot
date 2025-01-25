package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	r := gin.Default()
	initAuthRoutes(r)
	initBusinessRoutes(r)
	r.Run()
}
