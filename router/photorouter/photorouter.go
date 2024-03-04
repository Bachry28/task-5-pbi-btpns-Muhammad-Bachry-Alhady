package photorouter

import (
	"github.com/Bachry28/task-5-pbi-btpns-Muhammad-Bachry-Alhady/controllers/photocontroller"
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
func PhotoRouter(r *gin.Engine) {

	r.GET("/api/photo", helpers.Auth, photocontroller.GetAllPhoto)
	r.GET("/api/photo/:id", helpers.Auth, photocontroller.GetPhotoById)
	r.POST("/api/photo", helpers.Auth, photocontroller.CreatePhoto)
	r.PUT("/api/photo/:id", helpers.Auth, photocontroller.UpdatePhoto)
	r.DELETE("/api/photo/:id", helpers.Auth, photocontroller.DeletePhoto)

}
