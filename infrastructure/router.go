package infrastructure

import (
	gin "gopkg.in/gin-gonic/gin.v1"

	"github.com/yama-koo/go-api-example/interface/controllers"
)

// Router *gin.Engine
var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewSQLHandler())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })
	router.GET("/test", func(c *gin.Context) { c.JSON(200, map[string]interface{}{"text": "hogehoge"}) })

	Router = router
}
