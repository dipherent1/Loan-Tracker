package routers

import (
	"loaner/Deliverables/controllers"
	authmiddleware "loaner/Infrastructure/authMiddleware"
	repositories "loaner/Repositories"
	usecases "loaner/Usecases"
)

func UserRouter() {
	userRouter := Router.Group("/user")
	{
			// generate new auth repo
			userrepo := repositories.NewUserRepository(DataBase)
			userusecase := usecases.NewUserUsecase(userrepo)
			usercontroller := controllers.NewUserController(userusecase)

		// 	userRouter.POST("/register", usercontroller.Register)
		// 	userRouter.POST("/login", usercontroller.Login)
		userRouter.GET("/profile", authmiddleware.AuthMiddleware(), usercontroller.Profile)
	}
}
