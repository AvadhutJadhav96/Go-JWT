package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/AvadhutJadhav96/Go-JWT/controllers"              
)

func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
	//here token for validation is not requried since we have not signed up yet 
}