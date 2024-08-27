package routers

import (
	"fmt"
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

	fmt.Println("Router")
	fmt.Println(Router)
	fmt.Println("Router")
	// Initialize the Auth routes
	Router.POST("/register", authController.Register)
	// initialize the login route
	Router.POST("/login", authController.Login)
	// initialize the activate route
	Router.GET("/activate/:token", authController.Activate)

	
	
}