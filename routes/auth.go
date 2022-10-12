package routes

import (
	"blitzomni.com/m/controllers"
	"blitzomni.com/m/middleware"
	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
	authController controllers.AuthController
}

func NewAuthRoute(authController controllers.AuthController) AuthRoute {
	return AuthRoute{authController}
}

func (ar *AuthRoute) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/register", ar.authController.SignUp)
	router.POST("/login", ar.authController.Login)
	router.GET("/logout", middleware.Auth(), ar.authController.Logout)
	router.GET("/verify/:verificationCode", ar.authController.VerifyEmail)
}
