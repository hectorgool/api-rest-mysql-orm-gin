package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hectorgool/api-rest-mysql-orm-gin/config"
	"github.com/hectorgool/api-rest-mysql-orm-gin/model"
)

func DeletePerson(c *gin.Context) {

	id := c.Params.ByName("id")
	var person model.Person
	d := config.Dbh.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"data": id})

}

func UpdatePerson(c *gin.Context) {

	var person model.Person
	id := c.Params.ByName("id")
	if err := config.Dbh.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)
	config.Dbh.Save(&person)
	c.JSON(200, gin.H{"data": person})

}

func CreatePerson(c *gin.Context) {

	var person model.Person
	c.BindJSON(&person)
	config.Dbh.Create(&person)
	c.JSON(200, gin.H{"data": person})

}

func GetPerson(c *gin.Context) {

	id := c.Params.ByName("id")
	var person model.Person
	if err := config.Dbh.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{"data": person})
	}

}

func GetPersons(c *gin.Context) {

	var people []model.Person
	if err := config.Dbh.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{"data": people})
	}

}
