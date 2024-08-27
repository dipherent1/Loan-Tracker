package routers

import (
	"loaner/Deliverables/controllers"
	authmiddleware "loaner/Infrastructure/authMiddleware"
	repositories "loaner/Repositories"
	usecases "loaner/Usecases"
)

// refreshRouter
func RefreshTokenRouter() {
	refreshRouter := router.Group("/refresh")
	{
		// generate new auth repo
		refreshrepo := repositories.NewRefreshRepository(DataBase)
		refreshusecase := usecases.NewRefreshUseCase(refreshrepo)
		refreshcontroller := controllers.NewRefreshController(refreshusecase)

		refreshRouter.GET("", authmiddleware.AuthMiddleware(), refreshcontroller.Refresh)
	}
}