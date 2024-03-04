package userrouter

import (
	"github.com/Bachry28/task-5-pbi-btpns-Muhammad-Bachry-Alhady/controllers/usercontroller"
	"github.com/Bachry28/task-5-pbi-btpns-Muhammad-Bachry-Alhady/database"
	"github.com/Bachry28/task-5-pbi-btpns-Muhammad-Bachry-Alhady/helpers"

	"github.com/gin-gonic/gin"
)

func init() {
	// Load environment variables
	helpers.LoadEnv()

	// Connect to the database
	database.ConnectDatabase()
}

// UserRouter sets up routes for user-related endpoints
func UserRouter(r *gin.Engine) {

	r.GET("/api/users", usercontroller.GetAllUser)
	r.GET("/api/users/:id", usercontroller.GetUserById)
	r.POST("/api/users/register", usercontroller.Register)
	r.POST("/api/users/login", usercontroller.Login)
	r.PUT("/api/users/:id", usercontroller.UpdateUser)
	r.DELETE("/api/users/:id", usercontroller.DeleteUser)
}
