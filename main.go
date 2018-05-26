package main

import (
 	"github.com/gin-gonic/gin"
 	"github.com/hectorgool/gin2/config"
 	"github.com/hectorgool/gin2/controller"
)

func main() {

	defer config.Dbh.Close()

 	r := gin.Default()
 	r.Use(config.CORSMiddleware())

 	r.GET("/people/", controller.GetPersons)
 	r.GET("/people/:id", controller.GetPerson)
 	r.POST("/people", controller.CreatePerson)
 	r.PUT("/people/:id", controller.UpdatePerson)
 	r.DELETE("/people/:id", controller.DeletePerson)
 	
 	r.Run(":8080")

}

