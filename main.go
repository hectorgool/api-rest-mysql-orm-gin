package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hectorgool/api-rest-mysql-orm-gin/config"
	"github.com/hectorgool/api-rest-mysql-orm-gin/controller"
	"github.com/hectorgool/api-rest-mysql-orm-gin/common"
)

func main() {

	defer config.Dbh.Close()

	r := gin.Default()
	r.Use(common.CORSMiddleware())

	r.GET("/people/", controller.GetPersons)
	r.GET("/people/:id", controller.GetPerson)
	r.POST("/people", controller.CreatePerson)
	r.PUT("/people/:id", controller.UpdatePerson)
	r.DELETE("/people/:id", controller.DeletePerson)

	r.Run(":8080")

}
