package routers

import (
	"loaner/Deliverables/controllers"
	authmiddleware "loaner/Infrastructure/authMiddleware"
	repositories "loaner/Repositories"
	usecases "loaner/Usecases"
)

func LoanRouter() {
	loanRouter := Router.Group("/loan", authmiddleware.AuthMiddleware())
	{
		// generate repository
		loanRepo := repositories.NewLoanRepository(DataBase)
		loanUsecase := usecases.NewLoanUsecase(loanRepo)
		loanController := controllers.NewLoanController(loanUsecase)
		
		loanRouter.POST("/apply", loanController.Apply)
		loanRouter.GET("/get/:loanID", loanController.GetLoanById)
		// loanRouter.GET("/get", authmiddleware.AuthMiddleware(), loancontroller.Get)
		// loanRouter.GET("/getall", authmiddleware.AuthMiddleware(), loancontroller.GetAll)
		// loanRouter.POST("/approve", authmiddleware.AuthMiddleware(), loancontroller.Approve)
		// loanRouter.POST("/reject", authmiddleware.AuthMiddleware(), loancontroller.Reject)
		// loanRouter.POST("/pay", authmiddleware.AuthMiddleware(), loancontroller.Pay)
	}
}
