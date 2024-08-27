package routers

import (
	"loaner/Deliverables/controllers"
	repositories "loaner/Repositories"
	usecases "loaner/Usecases"
)


func AuthRouter () {
	// generate a new authrepo
	authRepo := repositories.NewAuthRepo(DataBase)
	// gerate a new authusecase
	authUsecase := usecases.NewAuthUsecase(authRepo)
	// generate a new authController
	authController := controllers.NewAuthController(authUsecase)

	// Initialize the Auth routes
	router.POST("/register", authController.Register)
	// initialize the login route
	router.POST("/login", authController.Login)
	// initialize the activate route
	router.GET("/activate/:token", authController.Activate)

	
	
}