package photocontroller

import (
	"net/http"

	"github.com/Bachry28/task-5-pbi-btpns-Muhammad-Bachry-Alhady/database"
	"github.com/Bachry28/task-5-pbi-btpns-Muhammad-Bachry-Alhady/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	database.ConnectDatabase()
}

func GetAllPhoto(c *gin.Context) {
	var photos []models.Photo

	if err := database.Database.Find(&photos).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"photo": photos})
}

func GetPhotoById(c *gin.Context) {
	var photo models.Photo
	id := c.Param("id")

	if err := database.Database.First(&photo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Photo not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"photo": photo})
}

func CreatePhoto(c *gin.Context) {
	var photo models.Photo
	user_id, _ := c.Get("user_id")

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, ok := user_id.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	photo.UserID = userID

	if err := database.Database.Create(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Create Success!", "photo": photo})
}

func UpdatePhoto(c *gin.Context) {
	var photo models.Photo
	id := c.Param("id")
	user_id, _ := c.Get("user_id")

	if err := database.Database.First(&photo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Photo not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	userID, ok := user_id.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	if userID != photo.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// photo.UserId = userID

	if err := database.Database.Save(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update Success!", "photo": photo})
}

func DeletePhoto(c *gin.Context) {
	var photo models.Photo
	id := c.Param("id")
	user_id, _ := c.Get("user_id")

	if err := database.Database.First(&photo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Photo not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	userID, ok := user_id.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	if userID != photo.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := database.Database.Delete(&photo, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete Success!", "photo": photo})
}
