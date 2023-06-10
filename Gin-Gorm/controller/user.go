package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kala-111/gin-gorm/config"
	"github.com/kala-111/gin-gorm/models"
)

func GetUsers(c *gin.Context) {
users:=[]models.User{}
config.DB.find(&users)
c.JSON(200, "Hello World")
}
func CreateUser(c *gin.Context) {
var user models.User
c.BindJSON(&user)
config.DB.Create(&user)
c.JSON(200, &user)

}
func DeleteUser(c *gin.Context) {
var user models.User
config.DB.Where("id=?", c.Param("id"))Delete(&user)
c.BindJSON(&user)


}
func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id=?",c.param("id"))First(&user)
	c.BindJSON(&user)
	config.DB.save(&user)
	c.JSON(200, &user)

}
			
