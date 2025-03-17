package routes

import(
	"github.com/AvadhutJadhav96/Go-JWT/controllers" 
	"github.com/gin-gonic/gin"
	"github.com/AvadhutJadhav96/Go-JWT/middleware" 	             
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	//used to authenticate the token

	incomingRoutes.GET("/users", controllers.GetUSsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}