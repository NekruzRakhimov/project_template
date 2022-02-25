package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project_template/utils"
)

func RunAllRoutes() {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Статус код 500, при любых panic()
	r.Use(gin.Recovery())

	// Запуск роутов
	initAllRoutes(r)

	// Запуск сервера
	_ = r.Run(utils.AppSettings.AppParams.PortRun)
}

func initAllRoutes(r *gin.Engine) {
	r.GET("/ping", PingPong)

}

// PingPong Проверка
func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}
