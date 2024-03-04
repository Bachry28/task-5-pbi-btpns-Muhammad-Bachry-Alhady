package usercontroller

import (
	"net/http"

	"github.com/Bachry28/task-5-pbi-btpns-Muhammad-Bachry-Alhady/database"
	"github.com/Bachry28/task-5-pbi-btpns-Muhammad-Bachry-Alhady/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllUser(c *gin.Context) {
	var users []models.User

	if err := database.Database.Preload("Photos").Find(&users).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := database.Database.Preload("Photos").First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Create Success!", "user": user})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if database.Database.Model(&models.User{}).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can't Update User!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update Success!", "user": user})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	if err := database.Database.Delete(&models.User{}, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can't Delete User!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete Success!", "user": user})
}
